package config

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Host string `toml:"host"`
	User string `toml:"user"`
	Password string `toml:"password"`
	Port int64 `toml:"port"`
	Name string `toml:"name"`
}