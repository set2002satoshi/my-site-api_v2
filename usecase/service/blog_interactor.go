package service

import (
	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
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
