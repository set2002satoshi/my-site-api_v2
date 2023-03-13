package category

import (
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database"
	"github.com/set2002satoshi/my-site-api_v2/interfaces/database/config"
	"github.com/set2002satoshi/my-site-api_v2/models"
	"github.com/set2002satoshi/my-site-api_v2/pkg/module/dto/response"
	"github.com/set2002satoshi/my-site-api_v2/usecase/service"
)

type CategoryController struct {
	Interactor service.CategoryInteractor
}

func NewCategoryController(db config.DB) *CategoryController {
	return &CategoryController{
		Interactor: service.CategoryInteractor{
			DB:           &config.DBRepository{DB: db},
			CategoryRepo: &database.CategoryRepository{},
			// HistoryCategoryRepo: &database.HistoryCategoryRepository{},
		},
	}
}

func (cc *CategoryController) convertActiveCategoryToDTO(obj *models.ActiveCategoryModel) response.ActiveCategoryEntity {
	return response.ActiveCategoryEntity{
		Id:           int(obj.GetCategoryId()),
		CategoryName: obj.GetCategoryName(),
	}
}
