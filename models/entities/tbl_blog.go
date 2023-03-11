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

// joinのクエリー用のstructです。 ※ マイグレーションしないように注意
type BlogWithNicknameEntity struct {
	BlogId    int       `gorm:"-:migration"`
	UserId    int       `gorm:"-:migration"`
	Nickname  string    `gorm:"-:migration"`
	Title     string    `gorm:"-:migration"`
	Context   string    `gorm:"-:migration"`
	Revision  int       `gorm:"-:migration"`
	CreatedAt time.Time `gorm:"-:migration"`
	UpdatedAt time.Time `gorm:"-:migration"`
}
