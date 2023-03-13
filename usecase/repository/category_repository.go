package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(tx *gorm.DB, obj *models.ActiveCategoryModel) (*models.ActiveCategoryModel, error)
}

type HistoryCategoryRepository interface {
}
