package impl

import (
	"github.com/sword-demon/go18/book/v4/apps/book"
)

type BookServiceImpl struct {
}

func NewBookServiceImpl() *BookServiceImpl {
	return &BookServiceImpl{}
}

// &BookServiceImpl 的空指针
// 强制必须实现接口的所有方法
var _ book.Service = (*BookServiceImpl)(nil)
