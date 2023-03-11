package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type BlogRepository interface {
	Create(db *gorm.DB, obj *models.ActiveBlogModel) (*models.ActiveBlogModel, error)
}

type HistoryBlogRepository interface {

}
