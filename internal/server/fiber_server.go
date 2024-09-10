package server

import (
	"log"
	"n1h41/oflow/internal/delivery/http/route"
	"n1h41/oflow/internal/model"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct{}

func NewFiberServer() Server {
	return &FiberServer{}
}

func (f *FiberServer) Run() {
	app := fiber.New(fiber.Config{
		// INFO: Setup default fiber error response
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusBadRequest).JSON(model.GlobalErrorHandlerResp{
				Success: false,
				Message: err.Error(),
			})
		},
	})

	// INFO: Attach routes to handlers
	route.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
