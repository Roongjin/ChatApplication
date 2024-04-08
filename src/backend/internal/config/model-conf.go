package config

type App struct {
	Database Database `mapstructure:"database"`
}

type Database struct {
	Redis BaseConfig `mapstructure:"redis"`
}

type BaseConfig struct {
	DSN string `mapstructure:"dsn"`
}
