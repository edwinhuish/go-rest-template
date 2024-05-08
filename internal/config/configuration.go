package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Logger   LoggerConfig
}

// SetupDB initialize configuration
func Setup(configPath string) {
	var configuration *Config

	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	config = configuration
}

// GetConfig helps you to get configuration data
func GetConfig() *Config {
	return config
}
