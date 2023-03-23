package entities

import (
	"strconv"
	"time"

	"gorm.io/gorm"
)

type TBLBlogWithCategoriesEntity struct {
	CompositeId string `gorm:"primaryKey"`
	CategoryId  int    `gorm:"not null"`
	BlogId      int    `gorm:"not null"`
}

func (b *TBLBlogWithCategoriesEntity) BeforeCreate(tx *gorm.DB) (err error) {
	b.CompositeId = strconv.Itoa(b.CategoryId) + "_" + strconv.Itoa(b.BlogId)
	return nil
}

type BlogWithNicknameattachCategoriesEntity struct {
	BlogId     int                `gorm:"-:migration"`
	UserId     int                `gorm:"-:migration"`
	Nickname   string             `gorm:"-:migration"`
	Title      string             `gorm:"-:migration"`
	Context    string             `gorm:"-:migration"`
	Categories []CategoriesEntity `gorm:"-:migration"`
	Revision   int                `gorm:"-:migration"`
	CreatedAt  time.Time          `gorm:"-:migration"`
	UpdatedAt  time.Time          `gorm:"-:migration"`
}

type CategoriesEntity struct {
	Id           string `gorm:"-:migration"`
	CategoryName string `gorm:"-:migration"`
}
