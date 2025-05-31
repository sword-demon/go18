// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package ioc

var containers = []Container{
	Api,
	Controller,
	Config,
	Default,
}

func Init() {
	for _, c := range containers {
		if err := c.Init(); err != nil {
			panic(err)
		}
	}
}

type Object interface {
	Init() error
}

type Container interface {
	Registry(name string, obj Object)
	Get(name string) Object
	// Init 初始化所有已经注册过的对象
	Init() error
}

type ObjectImpl struct {
}

func (o *ObjectImpl) Init() error {
	return nil
}
