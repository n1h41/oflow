package aws

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iot"
)

func TestInit(t *testing.T) {
	GetCognitoIdentityProviderClient()
}

func TestGetIotClient(t *testing.T) {
	client, err := GetIotClient()
	if err != nil {
		t.Error(err)
	}
	output, err := client.AttachPolicy(context.TODO(), &iot.AttachPolicyInput{
		PolicyName: aws.String("esp_p"),
		Target:     aws.String("us-east-1:a8eee6f0-8308-c13d-6231-9ca5eb663aae"),
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(output)
}
