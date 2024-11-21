package repository

import (
	"context"
	"n1h41/oflow/config"
	awsConfig "n1h41/oflow/internal/infrastructure/aws"
	"n1h41/oflow/internal/model"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentity"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/iot"
	"github.com/aws/aws-sdk-go-v2/service/iotdataplane"
	"github.com/awsdocs/aws-doc-sdk-examples/gov2/testtools"
)

func TestSignInUser(t *testing.T) {
	stubber := testtools.NewStubber()

	congnitoIdentityProvider := cognitoidentityprovider.NewFromConfig(*stubber.SdkConfig)
	cognitoIdentity := cognitoidentity.NewFromConfig(*stubber.SdkConfig)
	dynamoDB := dynamodb.NewFromConfig(*stubber.SdkConfig)
	iotClient := iot.NewFromConfig(*stubber.SdkConfig)
	userRepo := NewUserRepo(congnitoIdentityProvider, cognitoIdentity, dynamoDB, iotClient, "test", "test12345")

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

func TestUserLogin(t *testing.T) {
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

	config := config.Setup()
	userRepo := NewUserRepo(congnitoIdentityProvider, cognitoIdentity, nil, nil, config.AWS.ClientId, config.AWS.ClientSecret)

	loginCreds, err := userRepo.LoginUser(&model.SignInUserReq{
		Email:    "nihalninu25@gmail.com",
		Password: "nihal@23ktu",
	}, context.TODO())
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}

	t.Log(*loginCreds)
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

	config := config.Setup()
	userRepo := NewUserRepo(congnitoIdentityProvider, cognitoIdentity, nil, nil, config.AWS.ClientId, config.AWS.ClientSecret)

	userCreds, err := userRepo.FetchIdentityCredentials(context.TODO(), "eyJraWQiOiJDKzJVRlN0andSYUl0XC8rUEl6SEFvQjljcHp1MmRhdzYzNWo5RWVPT2hoUT0iLCJhbGciOiJSUzI1NiJ9.eyJzdWIiOiIxNGI4OTQ0OC1kMDExLTcwMDEtN2Q4MS03ZmUzMDg1ZTJhOTciLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLnVzLWVhc3QtMS5hbWF6b25hd3MuY29tXC91cy1lYXN0LTFfTThiOTdOd084IiwicGhvbmVfbnVtYmVyX3ZlcmlmaWVkIjpmYWxzZSwiY29nbml0bzp1c2VybmFtZSI6IjE0Yjg5NDQ4LWQwMTEtNzAwMS03ZDgxLTdmZTMwODVlMmE5NyIsIm9yaWdpbl9qdGkiOiI2YzliMDRhMi1kNzI4LTQ5ODUtYTA4MC1hNWNkNTIyOTM0OGUiLCJhdWQiOiI1dnJoOTN1NzBzdjhxZ2g5bjZmZzRyaDdsZSIsImN1c3RvbTpsYXN0X25hbWUiOiJBYmR1bGxhIiwiZXZlbnRfaWQiOiJmMGI3ZWRkYS1kODliLTQwYzItODA4NC1jYjNlMGM4YmYxYWYiLCJjdXN0b206Zmlyc3RfbmFtZSI6Ik5paGFsIiwidG9rZW5fdXNlIjoiaWQiLCJhdXRoX3RpbWUiOjE3MzA0NTEzMDksInBob25lX251bWJlciI6Iis5MTc1NTk4NjUzODYiLCJleHAiOjE3MzA0NTQ5MDksImlhdCI6MTczMDQ1MTMwOSwianRpIjoiNDg4ZTYwODMtMDQ1OS00NDc0LWFkNjUtMzc4MGJhZWU3NDgyIiwiZW1haWwiOiJuaWhhbG5pbnUyNUBnbWFpbC5jb20ifQ.KLRrfcZaGbhJrR85WVAJdWFDI0ZQnqQTzt4YGmhAP0pv3q_V-wg8G5TkIsvLVDA6lsJIjuE-xOsDhn-IrW5BjmrTk6ibg8SIj46lA9FP5N0puc1sLMZW2r8")
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(*userCreds.Credentials.AccessKeyId)
}

func TestFetchTopicLastMessage(t *testing.T) {
	iotDataPlaneClient, err := awsConfig.GetIotDataPlaneClien()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	retainedMessageOutput, err := iotDataPlaneClient.GetRetainedMessage(context.TODO(), &iotdataplane.GetRetainedMessageInput{
		Topic: aws.String("C4DEE2879A60/status"),
	})
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	encodedString := string(retainedMessageOutput.Payload)
	t.Log(encodedString)
}

func TestFetchDeviceList(t *testing.T) {
	dynamocDbClient, err := awsConfig.GetDynamoDbClient()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	scanOutput, err := dynamocDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String("devices"),
	})
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(scanOutput)
}

func TestAttachIotPolicyToIdentity(t *testing.T) {
	iotClient, err := awsConfig.GetIotClient()
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	output, err := iotClient.AttachPolicy(context.TODO(), &iot.AttachPolicyInput{
		PolicyName: aws.String("esp_p"),
		Target:     aws.String("us-east-1:a8eee6f0-8308-c13d-6231-9ca5eb663aae"),
	})
	if err != nil {
		t.Log(err)
		t.Fail()
		return
	}
	t.Log(output)
}
