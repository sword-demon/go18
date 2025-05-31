// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package generate

type StringSet struct {
	List []string `json:"list"`
}

type IntSet struct {
	List []int `json:"list"`
}

type Set[T any] struct {
	List []T `json:"list"`
}

func NewSet[T any]() *Set[T] {
	return &Set[T]{}
}

func (s *Set[T]) Add(item T) {
	s.List = append(s.List, item)
}
