// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package add_test

import (
	"github.com/go-playground/assert/v2"
	add "github.com/sword-demon/go18/skills/unit_test"
	"testing"
)

// TestAdd 针对 Add 函数的单元测试
func TestAdd(t *testing.T) {
	// -v -count=1
	t.Log(add.Add(1, 2))

	// 通过程序断言的方式来判断
	assert.Equal(t, 3, add.Add(1, 2))
	assert.NotEqual(t, 4, add.Add(1, 2))
}
