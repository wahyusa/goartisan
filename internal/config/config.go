package config

import (
	"bytes"
	_ "embed"

	"github.com/spf13/viper"
)

//go:embed config.toml
var defaultConfig []byte

type Config struct {
	Module struct {
		Name string
	}
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
	viper.AddConfigPath("$HOME/.goartisan")

	if err := viper.ReadInConfig(); err != nil {
		// If the config file is not found, use the embedded default configuration
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			viper.SetConfigType("toml")
			if err := viper.ReadConfig(bytes.NewBuffer(defaultConfig)); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
