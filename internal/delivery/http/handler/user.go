package handler

import (
	"log"
	"n1h41/oflow/internal/model"
	"n1h41/oflow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	SignUpUser(c *fiber.Ctx) error
}

type userHandler struct {
	userRepo repository.UserRepo
}

func NewUseHandler(userRepo repository.UserRepo) UserHandler {
	return &userHandler{
		userRepo: userRepo,
	}
}

func (h userHandler) SignUpUser(c *fiber.Ctx) error {
	var params model.SignUpUserReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	} // TODO:
	/* result, err := h.userRepo.SignUpUser(&params, c.Context())
	if err != nil {
		return err
	} */
	return c.JSON(fiber.Map{
		"message": params,
		"status":  true,
	})
}
