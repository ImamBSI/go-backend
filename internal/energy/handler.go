package energy

import "github.com/gofiber/fiber/v2"

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetRawEnergy(c *fiber.Ctx) error {
	year := c.Query("year")
	category := c.Query("category")

	result := h.service.GetRawEnergy(year, category)
	return c.JSON(result)
}

func (h *Handler) GetSumEnergy(c *fiber.Ctx) error {
	year := c.Query("year")
	category := c.Query("category")

	if category == "" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"error": "category is required"})
	}

	total := h.service.SumEnergy(year, category)

	return c.JSON(fiber.Map{
		"year":     year,
		"category": category,
		"total":    total,
	})
}
