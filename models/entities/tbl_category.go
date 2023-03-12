package entities

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type TBLCategoryEntity struct {
	CategoryID   types.IDENTIFICATION `gorm:"primaryKey"`
	CategoryName string               `gorm:"unique;not null;max:18"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
