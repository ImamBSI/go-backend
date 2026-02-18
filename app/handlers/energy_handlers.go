package handlers

import (
	"example.com/trial-go/app/models"
	"github.com/gofiber/fiber/v2"
)

var EnergyData []models.EnergyItem

func GetRawEnergy(c *fiber.Ctx) error {
	year := c.Query("year")
	category := c.Query("category")

	// 1. Filter berdasarkan Tahun (Jika ada parameter year)
	var filteredByYear []models.EnergyItem
	if year != "" {
		for _, item := range EnergyData {
			if item.Year == year {
				filteredByYear = append(filteredByYear, item)
			}
		}
	} else {
		// Jika tidak ada filter tahun, ambil semua data
		filteredByYear = EnergyData
	}

	// 2. Filter berdasarkan Category (Jika ada parameter category)
	// Kita gunakan []interface{} supaya bentuk JSON-nya fleksibel
	if category != "" {
		var finalResult []fiber.Map
		for _, item := range filteredByYear {
			val := 0.0
			switch category {
			case "electricity":
				val = item.Values.Electricity
			case "naturalGas":
				val = item.Values.NaturalGas
			case "productKl":
				val = item.Values.ProductKl
			case "indexEnergy":
				val = item.Values.IndexEnergy
			}

			finalResult = append(finalResult, fiber.Map{
				"month":    item.Month,
				"year":     item.Year,
				"category": category,
				"value":    val,
			})
		}
		return c.JSON(finalResult)
	}

	// 3. Default: Jika tidak ada category, kirim data mentah lengkap
	return c.JSON(filteredByYear)
}
