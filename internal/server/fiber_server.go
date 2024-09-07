package server

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/favicon"
)

type FiberServer struct{}

func NewFiberServer() Server {
	return &FiberServer{}
}

func (f *FiberServer) Run() {
	app := fiber.New()
	app.Use(favicon.New())
	// route.RegisterRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
