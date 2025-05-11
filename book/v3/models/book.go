package models

type Book struct {
	ID     uint    `json:"id" gorm:"primaryKey;column:id"`
	Title  string  `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author" gorm:"column:author;type:varchar(200)" validate:"required"`
	Price  float64 `json:"price" gorm:"column:price" validate:"required"`
	IsSale *bool   `json:"is_sale" gorm:"column:is_sale"`
}

func (b *Book) TableName() string {
	return "books"
}
