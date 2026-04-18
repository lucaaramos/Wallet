package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"wallet/services/accounts/internal/mongo"
	"wallet/services/accounts/internal/repository"
	"wallet/services/accounts/internal/services"
	"wallet/services/accounts/pkg/api/v1/accounts"
)

func main() {
	app := fiber.New()
	db := mongo.NewMongoClient()
	client := db.Database("accounts")
	// wiring
	repo := repository.NewAccountRepository(client)
	service := services.NewAccountService(repo)
	handler := accounts.NewHandler(service)

	accounts.RegisterRoutes(app, handler)

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	log.Println("running on 8080")
	app.Listen(":8080")
}
