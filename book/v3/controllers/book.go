// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package controllers

import (
	"context"
	"github.com/sword-demon/go18/book/v3/config"
	"github.com/sword-demon/go18/book/v3/models"
)

type BookController struct {
}

func NewBookController() *BookController {
	return &BookController{}
}

type (
	GetBookRequest struct {
		BookNumber string `json:"book_number"`
	}
)

func (c *BookController) GetBook(ctx context.Context, in *GetBookRequest) (*models.Book, error) {
	book := &models.Book{}
	if err := config.DB().Where("id = ?", in.BookNumber).Take(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}
