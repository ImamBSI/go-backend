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

// GetSumEnergy: default rekap per tahun, jika ada parameter maka filter
func (h *Handler) GetSumEnergy(c *fiber.Ctx) error {
	year := c.Query("year")
	category := c.Query("category")

	// Jika tidak ada parameter, tampilkan rekap per tahun
	if year == "" && category == "" {
		result := h.service.SumAllEnergyByYear()
		return c.JSON(result)
	}

	// Jika ada parameter, tampilkan total sesuai filter
	if category == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "category is required if filtering"})
	}
	total := h.service.SumEnergy(year, category)
	return c.JSON(fiber.Map{
		"year":     year,
		"category": category,
		"total":    total,
	})
}
