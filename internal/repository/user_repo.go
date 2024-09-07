package repository

import (
	"n1h41/oflow/internal/model"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type UserRepo interface {
	CreateUser(*model.CreatUserModel) error
}

type userRepo struct {
	userIdentityPoolClient *cognitoidentityprovider.Client
}

func NewUserRepo(userIdentityPoolClient *cognitoidentityprovider.Client) UserRepo {
	return userRepo{
		userIdentityPoolClient: userIdentityPoolClient,
	}
}

func (u userRepo) CreateUser(*model.CreatUserModel) error {
	return nil
}
