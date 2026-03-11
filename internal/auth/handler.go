package auth

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	Service *Service
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	result, err := h.Service.Login(req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(result)
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}
	result, err := h.Service.Register(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(result)
}

func (h *Handler) GetUsers(c *fiber.Ctx) error {

	users, err := h.Service.GetUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to fetch users"})
	}

	return c.JSON(users)
}

func (h *Handler) DeleteUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "invalid id"})
	}

	err = h.Service.DeleteUser(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": "failed to delete user"})
	}

	return c.JSON(fiber.Map{
		"message": "user deleted",
	})
}
