// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package ioc

var Api = NewMapContainer("api")
var Controller = NewMapContainer("controller")
var Config = NewMapContainer("config")
var Default = NewMapContainer("default")

// MapContainer ioc 容器
type MapContainer struct {
	name    string
	storage map[string]Object
}

func (m *MapContainer) Registry(name string, obj Object) {
	m.storage[name] = obj
}

func (m *MapContainer) Get(name string) Object {
	return m.storage[name]
}

// Init 初始化所有已经注册过的对象
func (m *MapContainer) Init() error {
	for _, obj := range m.storage {
		if err := obj.Init(); err != nil {
			return err
		}
	}

	return nil
}

func NewMapContainer(name string) *MapContainer {
	return &MapContainer{name: name, storage: make(map[string]Object)}
}
