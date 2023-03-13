package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(db *gorm.DB) ([]*models.ActiveCategoryModel, error)
	FindById(db *gorm.DB, id int) (*models.ActiveCategoryModel, error)
	Create(tx *gorm.DB, obj *models.ActiveCategoryModel) (*models.ActiveCategoryModel, error)
	DeleteById(db *gorm.DB, id int) error
}

type HistoryCategoryRepository interface {
}
