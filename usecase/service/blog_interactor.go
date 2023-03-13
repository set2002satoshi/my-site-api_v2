package service

import (
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
	DB              usecase.DBRepository
	UserRepo        repo.UserRepository
	BlogRepo        repo.BlogRepository
	HistoryBlogRepo repo.HistoryBlogRepository
}

func (bi *BlogInteractor) FindById(ctx *gin.Context, id int) (*models.ActiveBlogModel, error) {
	db := bi.DB.Connect()
	return bi.BlogRepo.FindById(db, id)
}

func (bi *BlogInteractor) FindAll(ctx *gin.Context) ([]*models.ActiveBlogModel, error) {
	db := bi.DB.Connect()
	return bi.BlogRepo.FindAll(db)
}

func (bi *BlogInteractor) Register(ctx *gin.Context, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	tx := bi.DB.Begin()
	createdBlog, err := bi.BlogRepo.Create(tx, obj)
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
		int(createdBlog.GetAuditTrail().GetRevision()),
		createdBlog.GetAuditTrail().GetCreatedAt(),
		createdBlog.GetAuditTrail().GetUpdatedAt(),
	)
}

func (bi *BlogInteractor) Update(ctx *gin.Context, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
	tx := bi.DB.Begin()
	current, err := bi.BlogRepo.FindById(tx, int(obj.GetBlogId()))
	if err != nil {
		return new(models.ActiveBlogModel), err
	}
	if current.GetUserId() != obj.GetUserId() {
		return new(models.ActiveBlogModel), errors.Add(errors.NewCustomError(), errors.SE0006)
	}

	if err := current.GetAuditTrail().CountUpRevision(obj.GetAuditTrail().GetRevision()); err != nil {
		return new(models.ActiveBlogModel), err
	}

	if _, err := bi.createHistoryBlog(tx, current); err != nil {
		tx.Rollback()
		return new(models.ActiveBlogModel), err
	}
	newBlog, err := models.NewActiveBlogModel(
		int(current.GetBlogId()),
		int(current.GetUserId()),
		current.GetNickname(),
		obj.GetTitle(),
		obj.GetContext(),
		int(current.GetAuditTrail().GetRevision()),
		current.GetAuditTrail().GetCreatedAt(),
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
	if _, err := bi.createHistoryBlog(tx, currentBlog); err != nil {
		return err
	}
	// historyBlog, err := models.NewHistoryBlogModel(
	// 	types.INITIAL_ID,
	// 	int(currentBlog.GetBlogId()),
	// 	int(currentBlog.GetUserId()),
	// 	currentBlog.GetTitle(),
	// 	currentBlog.GetContext(),
	// 	int(currentBlog.GetAuditTrail().GetRevision()),
	// 	currentBlog.GetAuditTrail().GetCreatedAt(),
	// 	currentBlog.GetAuditTrail().GetUpdatedAt(),
	// )
	// if err != nil {
	// 	return errors.Add(errors.NewCustomError(), errors.SE0005)
	// }
	// _, err = bi.HistoryBlogRepo.Create(tx, historyBlog)
	// if err != nil {
	// 	tx.Rollback()
	// 	return err
	// }
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

func (bi *BlogInteractor) createHistoryBlog(tx *gorm.DB, activeObj *models.ActiveBlogModel) (*models.HistoryBlogModel, error) {
	historyBlog, err := models.NewHistoryBlogModel(
		types.INITIAL_ID,
		int(activeObj.GetBlogId()),
		int(activeObj.GetUserId()),
		activeObj.GetTitle(),
		activeObj.GetContext(),
		int(activeObj.GetAuditTrail().GetRevision()),
		activeObj.GetAuditTrail().GetCreatedAt(),
		activeObj.GetAuditTrail().GetUpdatedAt(),
	)
	if err != nil {
		return new(models.HistoryBlogModel), errors.Add(errors.NewCustomError(), errors.SE0005)
	}
	return bi.HistoryBlogRepo.Create(tx, historyBlog)
}
