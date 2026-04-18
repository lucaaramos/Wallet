package accounts

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App, h *Handler) {
	api := app.Group("/api/v1/accounts")

	api.Post("/", h.CreateAccount)
	api.Get("/:id", h.GetAccount)

	api.Post("/deposit", h.Deposit)
	api.Post("/withdraw", h.Withdraw)
}
