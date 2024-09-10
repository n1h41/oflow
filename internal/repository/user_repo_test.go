package repository

import (
	"context"
	"n1h41/oflow/internal/model"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func TestSignInUser(t *testing.T) {
	stubber := testtools.NewStubber()

	congnitoIdentityProvider := cognitoidentityprovider.NewFromConfig(*stubber.SdkConfig)

	userRepo := NewUserRepo(congnitoIdentityProvider, "test")

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
		t.Fail()
		return
	}
	t.Log(*result.Session)
}
