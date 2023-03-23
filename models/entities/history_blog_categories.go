package entities

type HistoryBlogWithCategoriesEntity struct {
	Id         int `gorm:"primaryKey"`
	ActiveId   string
	CategoryId int
	BlogId     int
}
