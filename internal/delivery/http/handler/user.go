package handler

import (
	"log"
	"n1h41/oflow/internal/model"
	"n1h41/oflow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	SignUpUser(c *fiber.Ctx) error
	ConfirmUser(c *fiber.Ctx) error
	SignInUser(c *fiber.Ctx) error
	FetchIdentityCredentials(c *fiber.Ctx) error
	AddDevice(c *fiber.Ctx) error
	ListUserDevices(c *fiber.Ctx) error
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
	}
	result, err := h.userRepo.SignUpUser(&params, c.Context())
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": result,
		"status":  true,
	})
}

func (h *userHandler) ConfirmUser(c *fiber.Ctx) error {
	var params model.ConfirmUserReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	}
	result, err := h.userRepo.ConfirmUser(&params, c.Context())
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"message": result,
		"status":  true,
	})
}

func (h *userHandler) SignInUser(c *fiber.Ctx) error {
	var params model.SignInUserReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	}
	result, err := h.userRepo.LoginUser(&params, c.Context())
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": result,
	})
}

func (h *userHandler) FetchIdentityCredentials(c *fiber.Ctx) error {
	var params model.FetchIdentityCredentialsReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	}
	result, err := h.userRepo.FetchIdentityCredentials(c.Context(), params.Token)
	if err != nil {
		return err
	}
	return c.JSON(fiber.Map{
		"status":  true,
		"message": result,
	})
}

func (h *userHandler) AddDevice(c *fiber.Ctx) error {
	var params model.AddDeviceReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	}
	log.Println(params)
	// result, err := h.userRepo.AddDevice(c.Context(), params)
	panic("unimplemented")
}

func (h *userHandler) ListUserDevices(c *fiber.Ctx) error {
	var params model.ListUserDevicesReq
	if err := c.BodyParser(&params); err != nil {
		log.Println(err)
		return err
	}
	log.Println(params)
	panic("unimplemented")
}
