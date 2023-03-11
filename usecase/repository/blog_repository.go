package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	FindAll(db *gorm.DB) ([]*models.ActiveBlogModel, error)
	Create(db *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error)
}

type HistoryBlogRepository interface {

}
