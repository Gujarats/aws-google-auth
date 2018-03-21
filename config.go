package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	ProfileParent string `viper:profileParent`
	Profile       string `viper:"profile"`
	Password      string `viper:password`
}

const (
	pathConfig = ".aws-google-auth"
)

func getConfig() *Config {
	viper.AddConfigPath("$HOME/" + pathConfig)
	viper.SetConfigName("config")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file must exist in ~/"+pathConfig+"./config.yaml: %s \n", err))
	}

	config := &Config{}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("Fatal error unmarhsal config struct : %s \n", err))
	}

	return config
}
