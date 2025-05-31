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
