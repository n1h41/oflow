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

	cognitoIdentityPoolClient, err := aws.GetCognitoIdentityProviderClient()
	if err != nil {
		panic(err)
	}

	cognitoIdentityClient, err := aws.GetCognitoIdentityClient()
	if err != nil {
		panic(err)
	}

	dynamoDBClient, err := aws.GetDynamoDbClient()

	userRepo := repository.NewUserRepo(
		cognitoIdentityPoolClient,
		cognitoIdentityClient,
		dynamoDBClient,
		config.AWS.ClientId,
		config.AWS.ClientSecret,
	)
	userHandler := handler.NewUseHandler(userRepo)

	authGroup := app.Group("/auth")
	authGroup.Post("/sign-up", userHandler.SignUpUser)
	authGroup.Post("/confirm-user", userHandler.ConfirmUser)
	authGroup.Post("/sign-in", userHandler.SignInUser)

	deviceGroup := app.Group("/device")
	deviceGroup.Post("/add", userHandler.AddDevice)
}
