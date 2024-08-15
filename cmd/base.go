package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var baseCmd = &cobra.Command{
	Use:   "goartisan",
	Short: "GoArtisan is a CLI tool for generating Go REST API boilerplate",
	Long:  `GoArtisan helps you quickly scaffold a Go REST API project with various components.`,
}

func Execute() error {
	return baseCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	baseCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config.toml)")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// Set default values
	viper.SetDefault("app.folder", "app")
}
