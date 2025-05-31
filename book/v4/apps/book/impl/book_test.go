// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package impl

import (
	"context"
	"github.com/sword-demon/go18/book/v4/apps/book"
	"testing"
)

var ctx = context.Background()
var svc = NewBookServiceImpl()

// TDD: 测试驱动开发

func TestCreateBook(t *testing.T) {
	in := book.NewCreateBookRequest()
	in.SetIsSale(false)
	in.Author = "无解的游戏"
	in.Title = "kubernetes权威指南"
	in.Price = 123.121

	createBook, err := svc.CreateBook(ctx, in)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(createBook)
}
