package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"wallet/shared/rabbit"

	"wallet/services/accounts/internal/mongo"
	"wallet/services/accounts/internal/repository"
	"wallet/services/accounts/internal/services"
	"wallet/services/accounts/pkg/api/v1/accounts"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("no .env file found")
	}

	app := fiber.New()

	rabbitURL := os.Getenv("RABBITMQ_URL")
	if rabbitURL == "" {
		log.Fatal("RABBITMQ_URL is not set")
	}

	err = rabbit.Init(rabbitURL)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	// Mongo
	db := mongo.NewMongoClient()
	client := db.Database("accounts")

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
