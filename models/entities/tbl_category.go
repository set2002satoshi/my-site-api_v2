package entities

import (
	"time"
)

type TBLCategoryEntity struct {
	CategoryId   int    `gorm:"primaryKey"`
	CategoryName string `gorm:"unique;not null;max:18"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
