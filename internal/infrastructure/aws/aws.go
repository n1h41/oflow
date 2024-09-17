package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func getConfig() (*aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func GetCognitoIdentityProviderClient() (*cognitoidentityprovider.Client, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	client := cognitoidentityprovider.NewFromConfig(*cfg)
	return client, nil
}

func GetCognitoIdentityClient() (*cognitoidentity.Client, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	client := cognitoidentity.NewFromConfig(*cfg)
	return client, nil
}
