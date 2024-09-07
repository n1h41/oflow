package repository

import "n1h41/oflow/internal/model"

type UserRepo interface {
	CreateUser(*model.CreatUserModel)
}

type userRepo struct{}

func NewUserRepo() UserRepo {
	return &userRepo{}
}

func (u *userRepo) CreateUser(*model.CreatUserModel) {
	panic("unimplemented create user")
}
