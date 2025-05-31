// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/infraboard/mcube/v2/types"
	"github.com/sword-demon/go18/book/v4/apps/book"
)

func (b *BookServiceImpl) CreateBook(ctx context.Context, in *book.CreateBookRequest) (*book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookServiceImpl) QueryBook(ctx context.Context, in *book.QueryBookRequest) (*types.Set[*book.Book], error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookServiceImpl) FindBook(ctx context.Context, in *book.FindBookRequest) (*book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookServiceImpl) UpdateBook(ctx context.Context, in *book.UpdateBookRequest) (*book.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BookServiceImpl) DeleteBook(ctx context.Context, in *book.DeleteBookRequest) error {
	//TODO implement me
	panic("implement me")
}
