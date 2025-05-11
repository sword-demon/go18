// Copyright 2025 wxvirus(无解的游戏). All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/sword-demon/go18. The professional
// version of this repository is https://github.com/sword-demon/go18.

package config

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
	"os"
)

// config 全局变量，通过函数对这个进行访问
var config *Config

func C() *Config {
	// 如果没有配置文件，读取默认配置
	if config == nil {
		config = Default()
	}

	return config
}

// LoadConfigFromYaml 加载配置从 文件yaml 里读取 config
// 把外部配置读到全局变量里面来
func LoadConfigFromYaml(configPath string) error {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	config = C()
	return yaml.Unmarshal(content, config)
}

// LoadConfigFromEnv 从环境变量读取配置
func LoadConfigFromEnv() error {
	config = C()

	return env.Parse(config)
}
