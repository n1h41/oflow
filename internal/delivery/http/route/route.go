package route

import (
	"n1h41/oflow/config"
	"n1h41/oflow/internal/delivery/http/handler"
	"n1h41/oflow/internal/infrastructure/aws"
	"n1h41/oflow/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	config := config.Setup()

	userIdentityPoolClient, err := aws.GetUserIdentityClient()
	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(userIdentityPoolClient, config.AWS.ClientId)
	userHandler := handler.NewUseHandler(userRepo)

	authGroup := app.Group("/auth")
	authGroup.Post("/sign-up", userHandler.SignUpUser)
}
