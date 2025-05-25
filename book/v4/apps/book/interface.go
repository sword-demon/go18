package book

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
)

type Service interface {
	CreateBook(ctx context.Context, request *CreateBookRequest) (*Book, error)
	QueryBook(ctx context.Context, request *QueryBookRequest) (*types.Set[*Book], error)
}

type CreateBookRequest struct {
	// type 用于要使用gorm 来自动创建和更新表的时候 才需要定义
	Title  string  `json:"title"  gorm:"column:title;type:varchar(200)" validate:"required"`
	Author string  `json:"author"  gorm:"column:author;type:varchar(200);index" validate:"required"`
	Price  float64 `json:"price"  gorm:"column:price" validate:"required"`
	// bool false
	// nil 是零值, false
	IsSale *bool `json:"is_sale"  gorm:"column:is_sale"`
}

type QueryBookRequest struct {
}
