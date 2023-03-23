package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type BlogWithCategoryRepository interface {
	FindById(db *gorm.DB, id int) (*models.ActiveBlogWithCategoryModel, error)
	FindsByBlogId(db *gorm.DB, id int) ([]*models.ActiveBlogWithCategoryModel, error)
	Create(db *gorm.DB, obj *models.ActiveBlogWithCategoryModel) (*models.ActiveBlogWithCategoryModel, error)
	CreateAll(db *gorm.DB, obj []*models.ActiveBlogWithCategoryModel) ([]*models.ActiveBlogWithCategoryModel, error)
	DeleteById(db *gorm.DB, id string) error
}

type HistoryBlogWithCategoryRepository interface {
	Create(db *gorm.DB, obj *models.HistoryBlogWithCategoryModel) (*models.HistoryBlogWithCategoryModel, error)
}
