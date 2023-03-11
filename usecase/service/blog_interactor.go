package service

import (
	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/errors"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
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

func (bi BlogInteractor) Register(ctx *gin.Context, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error) {
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

func (bi *BlogInteractor) DeleteById(ctx *gin.Context, id int) error {
	tx := bi.DB.Begin()
	currentBlog, err := bi.BlogRepo.FindById(tx, id)
	if err != nil {
		return err
	}
	historyBlog, err := models.NewHistoryBlogModel(
		types.INITIAL_ID,
		int(currentBlog.GetBlogId()),
		int(currentBlog.GetUserId()),
		currentBlog.GetTitle(),
		currentBlog.GetContext(),
		int(currentBlog.GetAuditTrail().GetRevision()),
		currentBlog.GetAuditTrail().GetCreatedAt(),
		currentBlog.GetAuditTrail().GetUpdatedAt(),
	)
	if err != nil {
		return errors.Add(errors.NewCustomError(), errors.SE0005)
	}
	_, err = bi.HistoryBlogRepo.Create(tx, historyBlog)
	if err != nil {
		tx.Rollback()
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
