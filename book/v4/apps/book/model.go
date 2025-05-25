package book

import "github.com/infraboard/mcube/v2/tools/pretty"

type Book struct {
	Id uint `json:"id" gorm:"primaryKey;column:id"`

	CreateBookRequest
}

func (b *Book) TableName() string {
	return "books"
}

func (b *Book) String() string {
	return pretty.ToJSON(b)
}
