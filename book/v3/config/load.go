package config

import (
	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
	"gorm.io/gorm"
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

func DB() *gorm.DB {
	return C().MySQL.GetDB()
}
