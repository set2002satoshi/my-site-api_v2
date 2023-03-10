package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	FindById(db *gorm.DB, id int) (*models.ActiveBlogModel, error)
	FindAll(db *gorm.DB) ([]*models.ActiveBlogModel, error)
	Create(db *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error)
	Update(db *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error)
	DeleteById(db *gorm.DB, id int) error
}

type HistoryBlogRepository interface {
	Create(db *gorm.DB, obj *models.HistoryBlogModel) (*models.HistoryBlogModel, error)
}
