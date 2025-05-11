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
