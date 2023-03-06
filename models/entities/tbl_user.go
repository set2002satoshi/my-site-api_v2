package entities

import "time"

type TBLUserEntity struct {
	UserId    int    `gorm:"primaryKey"`
	Nickname  string `gorm:"not null;max:32"`
	Email     string `gorm:"unique;not null"`
	Password  []byte `gorm:"not null;max:32"`
	IconURL   string `gorm:"not null;max:255"`
	ImageKey  string
	Roll      string `gorm:"not null"`
	Revision  int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
