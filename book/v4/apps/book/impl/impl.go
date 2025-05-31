// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

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
