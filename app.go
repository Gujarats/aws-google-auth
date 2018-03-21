package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Gujarats/ark"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/service/ssm"
)

const (
	// for storing to env variable
	AWS_ACCESS_KEY_ID     = "AWS_ACCESS_KEY_ID"
	AWS_SECRET_ACCESS_KEY = "AWS_SECRET_ACCESS_KEY"
	AWS_SESSION_TOKEN     = "AWS_SESSION_TOKEN"
)

func main() {
	config := getConfig()

	app := "aws-google-auth"
	arg1 := "-p"
	profile := config.Profile

	result := runCommand(app, arg1, profile)
	fmt.Println(result)

	accessKey, secretKey := getKeysFromParameterStore(config)

	exportVariable(config, secretKey, accessKey)
}

func getKeysFromParameterStore(config *Config) (string, string) {
	sess, err := CreateSessionWithProfile(config.Region, config.Profile)

	creds := stscreds.NewCredentials(sess, config.RoleName)
	svc := ssm.New(sess, &aws.Config{Credentials: creds})

	accessKey, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(config.AccessKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}
	secretKey, err := svc.GetParameter(&ssm.GetParameterInput{
		Name:           aws.String(config.SecretKey),
		WithDecryption: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	return *accessKey.Parameter.Value, *secretKey.Parameter.Value
}

func runCommand(cmdName string, arg ...string) string {
	cmd := exec.Command(cmdName, arg...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	var stdout bytes.Buffer

	cmd.Stdout = &stdout

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Failed to start command %s. %s\n", cmdName, err.Error())
		os.Exit(1)
	}

	return string(stdout.Bytes())
}

func exportVariable(config *Config, secretKey, accessKey string) error {
	if config.UseEnvVariable {
		err := ark.SetEnvVariableAWS(accessKey, secretKey, "")
		if err != nil {
			return err
		}
	}

	if config.UseGradleProperties {
		configGradle := make(map[string]string)
		configGradle[ark.AccessKey] = config.GradleAccessKey
		configGradle[ark.SecretKey] = config.GradleSecretKey

		err := ark.UpdateGradleProperties(configGradle, accessKey, secretKey)
		if err != nil {
			return err
		}
	}

	return nil
}
