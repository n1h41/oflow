package repository

import (
	"context"
	"encoding/base64"
	"n1h41/oflow/internal/model"
	"n1h41/oflow/internal/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	cipTypes "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	dbTypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type UserRepo interface {
	SignUpUser(*model.SignUpUserReq, context.Context) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmUser(*model.ConfirmUserReq, context.Context) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	LoginUser(*model.SignInUserReq, context.Context) (*cognitoidentityprovider.InitiateAuthOutput, error)
	FetchCredentials(context.Context, string) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
	FetchDeviceList(*model.ListUserDevicesReq, context.Context) (*dynamodb.ScanOutput, error)
	AddDevice(*model.AddDeviceReq, context.Context) (*dynamodb.PutItemOutput, error)
}

type userRepo struct {
	cognitoIdentityProvider *cognitoidentityprovider.Client
	cognitoIdentity         *cognitoidentity.Client
	dynamoDBClient          *dynamodb.Client
	clientId                string
	clientSecret            string
}

func NewUserRepo(
	cognitoIdentityPoolClient *cognitoidentityprovider.Client,
	cognitoIdentityClient *cognitoidentity.Client,
	dynamoDBClient *dynamodb.Client,
	clientId string,
	clientSecret string,
) UserRepo {
	return userRepo{
		cognitoIdentityProvider: cognitoIdentityPoolClient,
		cognitoIdentity:         cognitoIdentityClient,
		dynamoDBClient:          dynamoDBClient,
		clientId:                clientId,
		clientSecret:            clientSecret,
	}
}

func (repo userRepo) SignUpUser(reqParam *model.SignUpUserReq, ctx context.Context) (*cognitoidentityprovider.SignUpOutput, error) {
	var userAttributes []cipTypes.AttributeType
	userAttributes = append(userAttributes, cipTypes.AttributeType{
		Name:  aws.String("custom:first_name"),
		Value: &reqParam.FirstName,
	})
	userAttributes = append(userAttributes, cipTypes.AttributeType{
		Name:  aws.String("custom:last_name"),
		Value: &reqParam.LastName,
	})
	userAttributes = append(userAttributes, cipTypes.AttributeType{
		Name:  aws.String("phone_number"),
		Value: &reqParam.Phone,
	})
	hash := util.GenerateHmacSHA256Hash(reqParam.Email+repo.clientId, repo.clientSecret)
	encodedSecretHash := base64.StdEncoding.EncodeToString(hash)
	result, err := repo.cognitoIdentityProvider.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		Username:       &reqParam.Email,
		Password:       &reqParam.Password,
		UserAttributes: userAttributes,
		ClientId:       aws.String(repo.clientId),
		SecretHash:     &encodedSecretHash,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo userRepo) ConfirmUser(reqParam *model.ConfirmUserReq, ctx context.Context) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
	hash := util.GenerateHmacSHA256Hash(reqParam.Email+repo.clientId, repo.clientSecret)
	encodedSecretHash := base64.StdEncoding.EncodeToString(hash)
	result, err := repo.cognitoIdentityProvider.ConfirmSignUp(ctx, &cognitoidentityprovider.ConfirmSignUpInput{
		ClientId:         &repo.clientId,
		Username:         &reqParam.Email,
		ConfirmationCode: &reqParam.ConfirmationCode,
		SecretHash:       &encodedSecretHash,
	})
	if err != nil {
		return nil, err
	}
	return result, err
}

func (repo userRepo) LoginUser(reqParam *model.SignInUserReq, ctx context.Context) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	authParams := make(map[string]string)
	hash := util.GenerateHmacSHA256Hash(reqParam.Email+repo.clientId, repo.clientSecret)
	encodedSecretHash := base64.StdEncoding.EncodeToString(hash)
	authParams["USERNAME"] = reqParam.Email
	authParams["PASSWORD"] = reqParam.Password
	authParams["SECRET_HASH"] = encodedSecretHash
	result, err := repo.cognitoIdentityProvider.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       cipTypes.AuthFlowType("USER_PASSWORD_AUTH"),
		ClientId:       aws.String(repo.clientId),
		AuthParameters: authParams,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo userRepo) FetchCredentials(ctx context.Context, token string) (*cognitoidentity.GetCredentialsForIdentityOutput, error) {
	getIdOutput, err := repo.cognitoIdentity.GetId(ctx, &cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String("us-east-1:cbad3db1-4d11-4c9a-9971-208eba32d1e3"),
		AccountId:      aws.String("725562617987"),
		Logins: map[string]string{
			"cognito-idp.us-east-1.amazonaws.com/us-east-1_M8b97NwO8": token,
		},
	})
	if err != nil {
		return nil, err
	}
	credentialOutput, err := repo.cognitoIdentity.GetCredentialsForIdentity(ctx, &cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: getIdOutput.IdentityId,
		Logins: map[string]string{
			"cognito-idp.us-east-1.amazonaws.com/us-east-1_M8b97NwO8": token,
		},
	})
	if err != nil {
		return nil, err
	}
	return credentialOutput, err
}

func (repo userRepo) FetchDeviceList(reqParam *model.ListUserDevicesReq, ctx context.Context) (*dynamodb.ScanOutput, error) {
	scanOutput, err := repo.dynamoDBClient.Scan(ctx, &dynamodb.ScanInput{
		TableName: aws.String("devices"),
	})
	if err != nil {
		return nil, err
	}
	return scanOutput, nil
}

func (repo userRepo) AddDevice(reqParam *model.AddDeviceReq, ctx context.Context) (*dynamodb.PutItemOutput, error) {
	putItemOutput, err := repo.dynamoDBClient.PutItem(ctx, &dynamodb.PutItemInput{
		Item: map[string]dbTypes.AttributeValue{
			"device_mac": &dbTypes.AttributeValueMemberS{Value: reqParam.DeviceMAC},
		},
		TableName: aws.String("devices"),
	})
	if err != nil {
		return nil, err
	}
	return putItemOutput, nil
}
