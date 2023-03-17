package entities

type HistoryBlogWithCategoriesEntity struct {
	HistoryID  int `gorm:"primaryKey"`
	Id         int
	CategoryId int
	BlogId     int
}
