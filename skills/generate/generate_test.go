// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package generate_test

import (
	"github.com/sword-demon/go18/skills/generate"
	"testing"
)

func TestStringSet(t *testing.T) {
	set := generate.NewSet[string]()
	set.Add("a")
	t.Log(set)
}

func TestIntSet(t *testing.T) {
	set := generate.NewSet[int]()
	set.Add(1)
	set.Add(2)
	t.Log(set)
}
