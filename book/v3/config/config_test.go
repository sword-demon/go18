package config

import (
	"testing"
)

func TestLoadConfigFromYaml(t *testing.T) {
	err := LoadConfigFromYaml("../application.yaml")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(C())
}
