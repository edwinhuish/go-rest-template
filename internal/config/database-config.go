package config

type DatabaseConfig struct {
	Driver       string `mapstructure:"driver"`
	Dbname       string `mapstructure:"dbname"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	Host         string `mapstructure:"host"`
	Port         string `mapstructure:"port"`
	MaxLifetime  int    `mapstructure:"max_life_time"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
