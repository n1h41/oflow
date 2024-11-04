package server

import (
	"errors"
	"log"
	"n1h41/oflow/internal/delivery/http/route"
	"net/http"

	"github.com/aws/smithy-go"
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
			var ae smithy.APIError
			if errors.As(err, &ae) {
				if ae.ErrorCode() == "NotAuthorizedException" {
					return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
						"status": false,
						"message": fiber.Map{
							"code":    ae.ErrorCode(),
							"message": ae.ErrorMessage(),
						},
					})
				}
				return c.JSON(fiber.Map{
					"status": false,
					"message": fiber.Map{
						"code":    ae.ErrorCode(),
						"message": ae.ErrorMessage(),
					},
				})
			}
			return c.JSON(fiber.Map{
				"status":  false,
				"message": err,
			})

		},
	})

	// INFO: Attach routes to handlers
	route.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
