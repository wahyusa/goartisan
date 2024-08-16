package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Folder string
	}
	Database struct {
		Default string
	}
	Templates struct {
		Path string
	}
	Structure struct {
		Dirs  []string
		Files []string
	}
}

func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
