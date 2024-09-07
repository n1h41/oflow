package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

func getConfig() *aws.Config {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	return &cfg
}

func GetUserIdentityClient() *cognitoidentityprovider.Client {
	cfg := *getConfig()
	client := cognitoidentityprovider.NewFromConfig(cfg)
	return client
}
