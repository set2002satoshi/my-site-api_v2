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
			CategoryRepo: &database.ActiveCategoryRepository{},
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

func (cc *CategoryController) convertActiveCategoryToDTOs(obj []*models.ActiveCategoryModel) []response.ActiveCategoryEntity {
	CEs := make([]response.ActiveCategoryEntity, len(obj))
	for i, v := range obj {
		CEs[i] = cc.convertActiveCategoryToDTO(v)
	}
	return CEs
}
