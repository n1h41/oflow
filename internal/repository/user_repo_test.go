package repository

import (
	"context"
	awsConfig "n1h41/oflow/internal/infrastructure/aws"
	"n1h41/oflow/internal/model"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func TestSignInUser(t *testing.T) {
	stubber := testtools.NewStubber()

	congnitoIdentityProvider := cognitoidentityprovider.NewFromConfig(*stubber.SdkConfig)
	cognitoIdentity := cognitoidentity.NewFromConfig(*stubber.SdkConfig)

	userRepo := NewUserRepo(congnitoIdentityProvider, cognitoIdentity, "test", "test12345")

	authParams := make(map[string]string)
	authParams["USERNAME"] = "test@gmail.com"
	authParams["PASSWORD"] = "test12345"
	stubber.Add(testtools.Stub{
		OperationName: "InitiateAuth",

		Input: &cognitoidentityprovider.InitiateAuthInput{
			AuthFlow:       types.AuthFlowType("USER_PASSWORD_AUTH"),
			ClientId:       aws.String("test"),
			AuthParameters: authParams,
		},
		Output: &cognitoidentityprovider.InitiateAuthOutput{
			AuthenticationResult: &types.AuthenticationResultType{},
			Session:              aws.String("meowmeow"),
		},
	})

	params := model.SignInUserReq{
		Email:    "test@gmail.com",
		Password: "test12345",
	}

	result, err := userRepo.LoginUser(&params, context.TODO())
	if err != nil {
		// t.Log(err)
		t.Fail()
		return
	}
	t.Log(*result.Session)
}

func TestFetchCredentials(t *testing.T) {
	congnitoIdentityProvider, err := awsConfig.GetCognitoIdentityProviderClient()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	cognitoIdentity, err := awsConfig.GetCognitoIdentityClient()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	userRepo := NewUserRepo(congnitoIdentityProvider, cognitoIdentity, "test", "test12345")

	loginCreds, err := userRepo.LoginUser(&model.SignInUserReq{
		Email:    "nihalninu25@gmail.com",
		Password: "nihal@23ktu",
	}, context.TODO())
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	userCreds, err := userRepo.FetchCredentials(context.TODO(), *loginCreds.AuthenticationResult.IdToken)
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(*userCreds.Credentials.AccessKeyId)
}
