package config

func Default() *Config {
	return &Config{}
}

type Config struct {
	Application *application `yaml:"app" toml:"app" json:"app"`
	MySQL       *mySQL       `yaml:"mysql" toml:"mysql" json:"mysql"`
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
}
