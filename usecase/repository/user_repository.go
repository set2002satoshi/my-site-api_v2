package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(db *gorm.DB, id int) (*models.ActiveUserModel, error)
	Create(db *gorm.DB, user *models.ActiveUserModel) (*models.ActiveUserModel, error)
}
