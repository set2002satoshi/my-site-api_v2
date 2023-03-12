package entities

import (
	"time"

	"github.com/set2002satoshi/my-site-api_v2/pkg/module/customs/types"
)

type HistoryCategoryEntity struct {
	ID           types.IDENTIFICATION `gorm:"primaryKey"`
	CategoryID   types.IDENTIFICATION
	CategoryName string    `gorm:"unique;not null;max:18"`
	CreatedAt    time.Time // active　timeを格納する
	DeletedAt    time.Time // 履歴が作られた時間を表示格納する
}
