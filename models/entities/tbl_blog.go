package entities

import "time"

type TBLBlogEntity struct {
	BlogId    int `gorm:"primaryKey"`
	UserId    int
	Title     string `gorm:"not null;max:32"`
	Context   string `gorm:"not null;max:255"`
	Revision  int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
