package entities

type TBLBlogWithCategoriesEntity struct {
	Id         int `gorm:"primaryKey"`
	CategoryId int `gorm:"primaryKey"`
	BlogId     int `gorm:"primaryKey"`
}
