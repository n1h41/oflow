package repository

import (
	"context"
	"encoding/base64"
	"n1h41/oflow/internal/model"
	"n1h41/oflow/internal/util"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type UserRepo interface {
	SignUpUser(*model.SignUpUserReq, context.Context) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmUser(*model.ConfirmUserReq, context.Context) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	LoginUser(*model.SignInUserReq, context.Context) (*cognitoidentityprovider.InitiateAuthOutput, error)
	FetchCredentials(context.Context, string) (*cognitoidentity.GetCredentialsForIdentityOutput, error)
}

type userRepo struct {
	cognitoIdentityProvider *cognitoidentityprovider.Client
	cognitoIdentity         *cognitoidentity.Client
	clientId                string
	clientSecret            string
}

func NewUserRepo(
	cognitoIdentityPoolClient *cognitoidentityprovider.Client,
	cognitoIdentityClient *cognitoidentity.Client,
	clientId string,
	clientSecret string,
) UserRepo {
	return userRepo{
		cognitoIdentityProvider: cognitoIdentityPoolClient,
		cognitoIdentity:         cognitoIdentityClient,
		clientId:                clientId,
		clientSecret:            clientSecret,
	}
}

func (repo userRepo) SignUpUser(reqParam *model.SignUpUserReq, ctx context.Context) (*cognitoidentityprovider.SignUpOutput, error) {
	var userAttributes []types.AttributeType
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("custom:first_name"),
		Value: &reqParam.FirstName,
	})
	userAttributes = append(userAttributes, types.AttributeType{
		Name:  aws.String("custom:last_name"),
		Value: &reqParam.LastName,
	})
	userAttributes = append(userAttributes, types.AttributeType{
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
		AuthFlow:       types.AuthFlowType("USER_PASSWORD_AUTH"),
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
		IdentityPoolId: aws.String(""),
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
