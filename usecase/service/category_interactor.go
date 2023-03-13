package service

import (
	"github.com/gin-gonic/gin"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/usecase"
	repo "github.com/set2002satoshi/my-site-api_v2/usecase/repository"
)

type CategoryInteractor struct {
	DB           usecase.DBRepository
	CategoryRepo repo.CategoryRepository
	// HistoryCategoryRepo repo.HistoryCategoryRepository
}

func (ci *CategoryInteractor) FindAll(ctx *gin.Context) ([]*models.ActiveCategoryModel, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.FindAll(db)
}

func (ci *CategoryInteractor) FindById(ctx *gin.Context, id int) (*models.ActiveCategoryModel, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.FindById(db, id)
}

func (ci *CategoryInteractor) Register(ctx *gin.Context, obj *models.ActiveCategoryModel) (*models.ActiveCategoryModel, error) {
	db := ci.DB.Connect()
	return ci.CategoryRepo.Create(db, obj)
}
