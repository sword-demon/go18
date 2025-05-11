package config

import (
	"fmt"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"github.com/sword-demon/go18/book/v3/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

func Default() *Config {
	return &Config{}
}

type Config struct {
	Application *application `yaml:"app" toml:"app" json:"app"`
	MySQL       *mySQL       `yaml:"mysql" toml:"mysql" json:"mysql"`
}

func (c *Config) String() string {
	return pretty.ToJSON(c)
}

type application struct {
	Host string `toml:"host" yaml:"host" json:"host" env:"HOST"`
	Port int    `toml:"port" yaml:"port" json:"port" env:"PORT"`
}

type mySQL struct {
	Host     string `toml:"host" yaml:"host" json:"host" env:"DATASOURCE_HOST"`
	Port     int    `toml:"port" yaml:"port" json:"port" env:"DATASOURCE_PORT"`
	DB       string `toml:"database" yaml:"database" json:"database" env:"DATASOURCE_DB"`
	Username string `toml:"username" yaml:"username" json:"username" env:"DATASOURCE_USERNAME"`
	Password string `toml:"password" yaml:"password" json:"password" env:"DATASOURCE_PASSWORD"`
	Debug    bool   `toml:"debug" yaml:"debug" json:"debug" env:"DATASOURCE_DEBUG"`

	// gorm db 对象，只需要有一个
	db *gorm.DB
	// 互斥锁
	lock sync.Mutex
}

func (m *mySQL) GetDB() *gorm.DB {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.db == nil {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			m.Username,
			m.Password,
			m.Host,
			m.Port,
			m.DB,
		)
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// 开启 debug 模式
		db = db.Debug()
		err = db.AutoMigrate(&models.Book{})
		if err != nil {
			return nil
		}
		m.db = db
	}

	return m.db
}
