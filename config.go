package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Profile string `viper:"profile"`

	Region            string `viper:"region"`
	RoleName          string `viper:roleName`
	AwsConfigPath     string `viper:awsConfigPath`
	AwsCredentialPath string `viper:awsCredentialPath`

	// Key for getting the value from parameter store
	SecretKey string `viper:secretKey`
	AccessKey string `viper:accessKey`

	// congfig key value for gralde.properties
	GradleAccessKey string `viper:gradleAccessKey`
	GradleSecretKey string `viper:gradleSecretKey`

	UseGradleProperties bool `viper:useGradleProperties`
	UseEnvVariable      bool `viper:useEnvVariable`
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
