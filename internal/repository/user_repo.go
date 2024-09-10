package repository

import (
	"context"
	"n1h41/oflow/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserRepo interface {
	SignUpUser(*model.SignUpUserReq, context.Context) (*cognitoidentityprovider.SignUpOutput, error)
	LoginUser(*model.SignInUserReq, context.Context) (*cognitoidentityprovider.InitiateAuthOutput, error)
}

type userRepo struct {
	userIdentityPoolClient *cognitoidentityprovider.Client
	clientId               string
}

func NewUserRepo(userIdentityPoolClient *cognitoidentityprovider.Client, clientId string) UserRepo {
	return userRepo{
		userIdentityPoolClient: userIdentityPoolClient,
		clientId:               clientId,
	}
}

func (repo userRepo) SignUpUser(reqParam *model.SignUpUserReq, ctx context.Context) (*cognitoidentityprovider.SignUpOutput, error) {
	var userAttributes []types.AttributeType
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("Email"),
		Value: &reqParam.Email,
	})
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("First Name"),
		Value: &reqParam.FirstName,
	})
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("Last Name"),
		Value: &reqParam.LastName,
	})
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("Phone"),
		Value: &reqParam.Phone,
	})
	result, err := repo.userIdentityPoolClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		Username:       &reqParam.Email,
		Password:       &reqParam.Password,
		UserAttributes: userAttributes,
		ClientId:       aws.String(repo.clientId),
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (repo userRepo) LoginUser(params *model.SignInUserReq, ctx context.Context) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	authParams := make(map[string]string)
	authParams["USERNAME"] = params.Email
	authParams["PASSWORD"] = params.Password
	result, err := repo.userIdentityPoolClient.InitiateAuth(ctx, &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       types.AuthFlowType("USER_PASSWORD_AUTH"),
		ClientId:       aws.String(repo.clientId),
		AuthParameters: authParams,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
