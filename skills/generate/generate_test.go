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
