package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func CreateSessionWithProfile(region string, profile string) (*session.Session, error) {
	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(region),
			Credentials: credentials.NewCredentials(
				&credentials.SharedCredentialsProvider{
					Profile: profile,
				},
			),
		},
	)

	if err != nil {
		return nil, err
	}

	return sess, nil
}
