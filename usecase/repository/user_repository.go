package repository

import (
	"github.com/set2002satoshi/my-site-api_v2/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(db *gorm.DB, id int) (*models.ActiveUserModel, error)
	FindAll(db *gorm.DB) ([]*models.ActiveUserModel, error)
	Create(db *gorm.DB, user *models.ActiveUserModel) (*models.ActiveUserModel, error)
	Update(tx *gorm.DB, user *models.ActiveUserModel) (*models.ActiveUserModel, error)
	DeleteById(tx *gorm.DB, id int) error
}

type HistoryUserRepository interface {
	FindById(db *gorm.DB, id int) (*models.HistoryUserModel, error)
	Create(tx *gorm.DB, user *models.HistoryUserModel) (*models.HistoryUserModel, error)
	DeleteById(tx *gorm.DB, id int) error
}