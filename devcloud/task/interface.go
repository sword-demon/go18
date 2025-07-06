// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package task

import "context"

type Service interface {
	Run(context.Context) (*Task, error)
}

type Task struct {
	Id string `json:"id" gorm:"column:id;type:varchar(60)" description:"id"`
}

type TaskSpec struct {
}
