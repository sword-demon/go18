// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

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
