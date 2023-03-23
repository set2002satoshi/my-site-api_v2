package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
	"gorm.io/gorm"
)

type BlogInteractor struct {
	DB                          usecase.DBRepository
	UserRepo                    repo.UserRepository
	BlogRepo                    repo.BlogRepository
	HistoryBlogRepo             repo.HistoryBlogRepository
	BlogWithCategoryRepo        repo.BlogWithCategoryRepository
	HistoryBlogWithCategoryRepo repo.HistoryBlogWithCategoryRepository
	CategoryRepo                repo.CategoryRepository
	// HistoryCategoryRepo  repo.HistoryCategoryRepository
}

func (bi *BlogInteractor) FindById(ctx *gin.Context, id int) (*models.ActiveBlogModel, error) {
	db := bi.DB.Connect()
	blog, err := bi.BlogRepo.FindById(db, id)
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	blogWithCategoryIds, err := bi.BlogWithCategoryRepo.FindsByBlogId(db, int(blog.GetBlogId()))
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	categoryIds, err := bi.extractCategoryIds(blogWithCategoryIds)
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	categories, err := bi.CategoryRepo.FindByIds(db, categoryIds)
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	return models.NewActiveBlogModel(
		int(blog.GetBlogId()),
		int(blog.GetUserId()),
		blog.GetNickname(),
		blog.GetTitle(),
		blog.GetContext(),
		categories,
		int(blog.GetAuditTrail().GetRevision()),
		blog.GetAuditTrail().GetCreatedAt(),
		blog.GetAuditTrail().GetUpdatedAt(),
	)

}

func (bi *BlogInteractor) FindAll(ctx *gin.Context) ([]*models.ActiveBlogModel, error) {
	db := bi.DB.Connect()
	blog, err := bi.BlogRepo.FindAll(db)
	if err != nil {
		return make([]*models.ActiveBlogModel, 0), err
	}
	blogWithCategories := make([]*models.ActiveBlogModel, len(blog))
	for i, v := range blog {
		blogWithCategoryIds, err := bi.BlogWithCategoryRepo.FindsByBlogId(db, int(v.GetBlogId()))
		if err != nil {
			return make([]*models.ActiveBlogModel, 0), err
		}
		ids, err := bi.extractCategoryIds(blogWithCategoryIds)
		if err != nil {
			return make([]*models.ActiveBlogModel, 0), err
		}
		categories, err := bi.CategoryRepo.FindByIds(db, ids)
		if err != nil {
			return make([]*models.ActiveBlogModel, 0), err
		}
		blog, err := models.NewActiveBlogModel(
			int(v.GetBlogId()),
			int(v.GetUserId()),
			v.GetNickname(),
			v.GetTitle(),
			v.GetContext(),
			categories,
			int(v.GetAuditTrail().GetRevision()),
			v.GetAuditTrail().GetCreatedAt(),
			v.GetAuditTrail().GetUpdatedAt(),
		)
		if err != nil {
			return make([]*models.ActiveBlogModel, 0), err
		}
		blogWithCategories[i] = blog
	}
	return blogWithCategories, nil
}

func (bi *BlogInteractor) extractCategoryIds(blogWithCategoryIds []*models.ActiveBlogWithCategoryModel) ([]int, error) {
	numbers := make([]int, len(blogWithCategoryIds))
	for i, v := range blogWithCategoryIds {
		numbers[i] = int(v.GetCategoryId())
	}
	return numbers, nil
}

func (bi *BlogInteractor) Register(ctx *gin.Context, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	tx := bi.DB.Begin()

	createdBlog, err := bi.BlogRepo.Create(tx, obj)
	if err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}

	categories, err := bi.createCategories(tx, int(createdBlog.GetBlogId()), obj.GetCategoryIds())
	if err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}

	acquiredUser, err := bi.UserRepo.FindById(tx, int(createdBlog.GetUserId()))
	if err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}

	return models.NewActiveBlogModel(
		int(createdBlog.GetBlogId()),
		int(createdBlog.GetUserId()),
		acquiredUser.GetNickname(),
		createdBlog.GetTitle(),
		createdBlog.GetContext(),
		categories,
		int(createdBlog.GetAuditTrail().GetRevision()),
		createdBlog.GetAuditTrail().GetCreatedAt(),
		createdBlog.GetAuditTrail().GetUpdatedAt(),
	)
}

func (bi *BlogInteractor) Update(ctx *gin.Context, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	tx := bi.DB.Begin()
	currentBlog, err := bi.BlogRepo.FindById(tx, int(obj.GetBlogId()))
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	if currentBlog.GetUserId() != obj.GetUserId() {
		return new(models.ActiveBlogModel), errors.Add(errors.NewCustomError(), errors.SE0006)
	}

	if err := currentBlog.GetAuditTrail().CountUpRevision(obj.GetAuditTrail().GetRevision()); err != nil {
		return new(models.ActiveBlogModel), err
	}

	currentIds, err := bi.BlogWithCategoryRepo.FindsByBlogId(tx, int(currentBlog.GetBlogId()))
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	// blogとcategoryのhistoryを作成
	if _, err := bi.createHistoryBlogWithCategories(tx, currentBlog, currentIds); err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}
	for _, v := range currentIds {
		err = bi.BlogWithCategoryRepo.DeleteById(tx, v.GetId())
		if err != nil {
			return new(models.ActiveBlogModel), err
		}
	}

	// 元々あったカテゴリを削除
	// err = bi.BlogWithCategoryRepo.DeleteAll(tx, currentIds)
	// if err != nil {
	// 	tx.Rollback()
	// 	return new(models.ActiveBlogModel), err
	// }

	// 新たにカテゴリを作成
	categories, err := bi.createCategories(tx, int(currentBlog.GetBlogId()), obj.GetCategoryIds())
	if err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}
	fmt.Println(categories)

	// カテゴリの処理もかく
	newBlog, err := models.NewActiveBlogModel(
		int(currentBlog.GetBlogId()),
		int(currentBlog.GetUserId()),
		currentBlog.GetNickname(),
		obj.GetTitle(),
		obj.GetContext(),
		categories,
		int(currentBlog.GetAuditTrail().GetRevision()),
		currentBlog.GetAuditTrail().GetCreatedAt(),
		time.Now(),
	)
	if err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}
	if _, err := bi.BlogRepo.Update(tx, newBlog); err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return &models.ActiveBlogModel{}, err
	}
	return newBlog, nil

}

func (bi *BlogInteractor) DeleteById(ctx *gin.Context, id int) error {
	tx := bi.DB.Begin()
	currentBlog, err := bi.BlogRepo.FindById(tx, id)
	if err != nil {
		return err
	}
	currentIds, err := bi.BlogWithCategoryRepo.FindsByBlogId(tx, int(currentBlog.GetBlogId()))
	if err != nil {
		return err
	}
	if _, err := bi.createHistoryBlogWithCategories(tx, currentBlog, currentIds); err != nil {
		return err
	}
	err = bi.BlogRepo.DeleteById(tx, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (bi *BlogInteractor) createHistoryBlogWithCategories(tx *gorm.DB, activeObj *models.ActiveBlogModel, activeCategories []*models.ActiveBlogWithCategoryModel) (*models.HistoryBlogModel, error) {
	bcm := make([]*models.HistoryBlogWithCategoryModel, len(activeCategories))
	for i, v := range activeCategories {
		historyCategory, err := models.NewHistoryBlogWithCategoryModel(
			types.INITIAL_ID,
			v.GetId(),
			int(v.GetCategoryId()),
			int(v.GetBlogId()),
		)
		if err != nil {
			return new(models.HistoryBlogModel), errors.Add(errors.NewCustomError(), errors.SE0005)
		}
		CreatedHistoryCategory, err := bi.HistoryBlogWithCategoryRepo.Create(tx, historyCategory)
		if err != nil {
			return new(models.HistoryBlogModel), err
		}
		bcm[i] = CreatedHistoryCategory
	}

	historyBlog, err := models.NewHistoryBlogModel(
		types.INITIAL_ID,
		int(activeObj.GetBlogId()),
		int(activeObj.GetUserId()),
		activeObj.GetTitle(),
		activeObj.GetContext(),
		[]*models.HistoryCategoryModel{}, // めんどうくさいからからのストラクトを代入
		int(activeObj.GetAuditTrail().GetRevision()),
		activeObj.GetAuditTrail().GetCreatedAt(),
		activeObj.GetAuditTrail().GetUpdatedAt(),
	)
	if err != nil {
		return new(models.HistoryBlogModel), errors.Add(errors.NewCustomError(), errors.SE0005)
	}
	return bi.HistoryBlogRepo.Create(tx, historyBlog)
}

func (bi *BlogInteractor) createCategories(tx *gorm.DB, blogId int, categoriesId []*models.ActiveCategoryModel) ([]*models.ActiveCategoryModel, error) {
	fmt.Println(categoriesId)
	bwc := make([]*models.ActiveBlogWithCategoryModel, len(categoriesId))
	for i, category := range categoriesId {
		blogWithCategoryIds, err := models.NewActiveBlogWithCategoryModel(
			"",
			int(category.GetCategoryId()),
			blogId,
		)
		if err != nil {
			return make([]*models.ActiveCategoryModel, 0), err
		}
		_, err = bi.BlogWithCategoryRepo.Create(tx, blogWithCategoryIds) // 一個一個作成
		if err != nil {
			return make([]*models.ActiveCategoryModel, 0), err
		}
		bwc[i] = blogWithCategoryIds
	}
	// _, err := bi.BlogWithCategoryRepo.CreateAll(tx, bwc)

	cm := make([]int, len(categoriesId))
	for i, v := range bwc {
		cm[i] = int(v.GetCategoryId())
	}
	return bi.CategoryRepo.FindByIds(tx, cm)
}
