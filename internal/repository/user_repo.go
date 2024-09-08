package repository

import (
	"context"
	"n1h41/oflow/internal/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserRepo interface {
	SignUpUser(*model.CreateUserModelReq, context.Context) (*cognitoidentityprovider.SignUpOutput, error)
}

type userRepo struct {
	userIdentityPoolClient *cognitoidentityprovider.Client
}

func NewUserRepo(userIdentityPoolClient *cognitoidentityprovider.Client) UserRepo {
	return userRepo{
		userIdentityPoolClient: userIdentityPoolClient,
	}
}

func (repo userRepo) SignUpUser(reqParam *model.CreateUserModelReq, ctx context.Context) (*cognitoidentityprovider.SignUpOutput, error) {
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
	// TODO: Need to add the ID of the client associated with the user pool
	result, err := repo.userIdentityPoolClient.SignUp(ctx, &cognitoidentityprovider.SignUpInput{
		Username:       &reqParam.Email,
		Password:       &reqParam.Password,
		UserAttributes: userAttributes,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
