package config

type App struct {
	Database Database `mapstructure:"database"`
	Network  Network  `mapstructure:"network"`
}

type Database struct {
	Redis    BaseConfig `mapstructure:"redis"`
	Postgres BaseConfig `mapstructure:"postgres"`
}

type BaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

type BaseIpConfig struct {
	IpAddr string `mapstructure:"ip_address"`
}

type Network struct {
	En0 BaseIpConfig `mapstructure:"en0"`
}
