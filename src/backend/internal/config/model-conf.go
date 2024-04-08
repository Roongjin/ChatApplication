package config

type App struct {
	Database Database `mapstructure:"database"`
}

type Database struct {
	Redis    BaseConfig `mapstructure:"redis"`
	Postgres BaseConfig `mapstructure:"postgres"`
}

type BaseConfig struct {
	DSN string `mapstructure:"dsn"`
}
