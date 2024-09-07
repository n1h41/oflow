package handler

import (
	"n1h41/oflow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	CreateUser(ctx *fiber.Ctx) error
}

type userHandler struct {
	userRepo *repository.UserRepo
}

func NewUseHandler(userRepo *repository.UserRepo) UserHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (h *userHandler) CreateUser(ctx *fiber.Ctx) error {
	panic("unimplemented Create User handler")
}
