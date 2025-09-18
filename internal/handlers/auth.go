package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/wansanjou/poke-api/internal/core/domains"
	"github.com/wansanjou/poke-api/internal/core/ports"
)

type authHandler struct {
	authService ports.AuthService
}

func NewAuthHandler(authService ports.AuthService) *authHandler {
	return &authHandler{
		authService,
	}
}

func (h *authHandler) Register(c *fiber.Ctx) error {
	var req domains.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	ctx := context.Background()
	user, err := h.authService.Register(ctx, domains.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *authHandler) Login(c *fiber.Ctx) error {
	var req domains.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	ctx := context.Background()
	resp, err := h.authService.Login(ctx, domains.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": resp.Token,
	})
}
