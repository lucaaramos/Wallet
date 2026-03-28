package accounts

import (
	"wallet/services/accounts/internal/models"

	"github.com/gofiber/fiber/v2"
)

type AccountService interface {
	CreateAccount(userID string) (*models.Account, error)
	GetAccount(id string) (*models.Account, error)
	Deposit(id string, amount float64) error
	Withdraw(id string, amount float64) error
}

type Handler struct {
	service AccountService
}

func NewHandler(s AccountService) *Handler {
	return &Handler{service: s}
}

func (h *Handler) CreateAccount(c *fiber.Ctx) error {
	var req struct {
		UserID string `json:"user_id"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err)
	}

	acc, err := h.service.CreateAccount(req.UserID)
	if err != nil {
		return c.Status(500).JSON(err)
	}

	return c.JSON(acc)
}

func (h *Handler) GetAccount(c *fiber.Ctx) error {
	id := c.Params("id")

	acc, err := h.service.GetAccount(id)
	if err != nil {
		return c.Status(404).JSON(err)
	}

	return c.JSON(acc)
}

func (h *Handler) Deposit(c *fiber.Ctx) error {
	var req struct {
		AccountID string  `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err)
	}

	err := h.service.Deposit(req.AccountID, req.Amount)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON("ok")
}

func (h *Handler) Withdraw(c *fiber.Ctx) error {
	var req struct {
		AccountID string  `json:"account_id"`
		Amount    float64 `json:"amount"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(err)
	}

	err := h.service.Withdraw(req.AccountID, req.Amount)
	if err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON("ok")
}
