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
