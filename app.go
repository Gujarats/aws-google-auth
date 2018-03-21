package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"

	"github.com/Gujarats/logger"
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
	profileParent := config.ProfileParent

	result := runCommand(app, arg1, profileParent)
	fmt.Println(result)

	assumeRole(config)
}

func assumeRole(config *Config) {

	app := "awsudo"
	argUser := config.Profile

	result := runCommand(app, argUser)

	lines := strings.Fields(result)
	fmt.Println("len lines = ", len(lines))
	fmt.Println("len 0 = ", lines[0])
	logger.Debug("lines", lines)

	exportVariable(lines)
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

func exportVariable(data []string) error {
	envs := make(map[string]string)
	envs[AWS_ACCESS_KEY_ID] = data[0][len(AWS_ACCESS_KEY_ID)+1:]
	envs[AWS_SECRET_ACCESS_KEY] = data[1][len(AWS_SECRET_ACCESS_KEY)+1:]
	envs[AWS_SESSION_TOKEN] = data[2][len(AWS_SESSION_TOKEN)+1:]

	for key, value := range envs {
		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}

	logger.Debug("AWS_ACCESS_KEY_ID :: ", os.Getenv(AWS_ACCESS_KEY_ID))
	logger.Debug("AWS_SECRET_ACCESS_KEY :: ", os.Getenv(AWS_SECRET_ACCESS_KEY))
	logger.Debug("AWS_SECRET_ACCESS_KEY :: ", os.Getenv(AWS_SESSION_TOKEN))

	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())

	return nil
}
