package server

import (
	"log"
	"n1h41/oflow/internal/delivery/http/route"

	"github.com/gofiber/fiber/v2"
)

type FiberServer struct{}

func NewFiberServer() Server {
	return &FiberServer{}
}

func (f *FiberServer) Run() {
	app := fiber.New()

	// INFO: Attach routes to handlers
	route.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
