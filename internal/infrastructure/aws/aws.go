package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
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

func GetDynamoDbClient() (*dynamodb.Client, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	client := dynamodb.NewFromConfig(*cfg)
	return client, nil
}

func GetIotDataPlaneClien() (*iotdataplane.Client, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	client := iotdataplane.NewFromConfig(*cfg)
	return client, nil
}

func GetIotClient() (*iot.Client, error) {
	cfg, err := getConfig()
	if err != nil {
		return nil, err
	}
	client := iot.NewFromConfig(*cfg)
	return client, nil
}
