package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Logger LoggerConfig
	HTTP   HTTPConfig
	Sender SenderConfig
}

type LoggerConfig struct {
	MinLevel int8
}
type HTTPConfig struct {
	Port int
}

type SenderConfig struct {
	IsPool     bool
	NumWorkers uint8
	Adapter    string
}

func loadConfig(configFilePath string) (*Config, error) {
	viper.SetConfigFile(configFilePath)

	setDefaults()

	if configFilePath != "" {
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func setDefaults() {
	viper.SetDefault("logger.minLevel", 0)

	viper.SetDefault("http.port", 8080)

	viper.SetDefault("sender.isPool", false)
	viper.SetDefault("sender.numWorkers", 3)
	viper.SetDefault("sender.adapter", "log")
}
