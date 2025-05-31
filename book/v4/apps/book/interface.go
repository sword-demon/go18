// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package book

import (
	"context"
	"github.com/infraboard/mcube/v2/http/request"
	"github.com/infraboard/mcube/v2/ioc/config/validator"
	"github.com/infraboard/mcube/v2/types"
)

// Service 接口的定义
type Service interface {
	CreateBook(ctx context.Context, in *CreateBookRequest) (*Book, error)
	QueryBook(ctx context.Context, in *QueryBookRequest) (*types.Set[*Book], error)
	FindBook(ctx context.Context, in *FindBookRequest) (*Book, error)
	UpdateBook(ctx context.Context, in *UpdateBookRequest) (*Book, error)
	DeleteBook(ctx context.Context, in *DeleteBookRequest) error
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

func NewCreateBookRequest() *CreateBookRequest {
	return &CreateBookRequest{}
}

func (r *CreateBookRequest) SetIsSale(v bool) *CreateBookRequest {
	r.IsSale = &v
	return r
}

func (r *CreateBookRequest) Validate() error {
	return validator.Validate(r)
}

type DeleteBookRequest struct {
	FindBookRequest
}

type UpdateBookRequest struct {
	FindBookRequest
	CreateBookRequest
}

type FindBookRequest struct {
	Id uint `json:"id"`
}

type QueryBookRequest struct {
	request.PageRequest

	Keywords string `json:"keywords"`
}

func NewQueryBookRequest() *QueryBookRequest {
	// 给分页请求参数传递一个默认值
	return &QueryBookRequest{PageRequest: *request.NewDefaultPageRequest()}
}
