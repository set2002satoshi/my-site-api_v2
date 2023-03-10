package entities

import "time"

type HistoryBlogEntity struct {
	Id        int `gorm:"primaryKey"`
	ActiveId  int
	UserId    int
	Title     string    `gorm:"not null;max:32"`
	Context   string    `gorm:"not null;max:255"`
	Revision  int       `gorm:"not null"`
	CreatedAt time.Time // active　timeを格納する
	DeletedAt time.Time // 履歴が作られた時間を表示格納する
}
