package main

import (
	"github.com/spf13/viper"
)

type Config struct {
	Http HttpConfig
}

type HttpConfig struct {
	Port int
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
	viper.SetDefault("http.port", 8080)
}
