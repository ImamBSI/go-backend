package services

import "example.com/trial-go/app/models"

// FilterRawEnergy menyaring data berdasarkan tahun dan category
func FilterRawEnergy(data []models.EnergyItem, year string, category string) []map[string]interface{} {
	var results []map[string]interface{}

	for _, item := range data {
		if item.Year == year {
			val := 0.0
			switch category {
			case "electricity":
				val = item.Values.Electricity
			case "naturalGas":
				val = item.Values.NaturalGas
			case "productKl":
				val = item.Values.ProductKl
			}

			results = append(results, map[string]interface{}{
				"month": item.Month,
				"value": val,
			})
		}
	}
	return results
}

// SumEnergy menghitung total berdasarkan kategori dan tahun
func SumEnergy(data []models.EnergyItem, year string, category string) float64 {
	var total float64
	for _, item := range data {
		// Jika year kosong, hitung semua tahun (default)
		if year == "" || item.Year == year {
			switch category {
			case "electricity":
				total += item.Values.Electricity
			case "naturalGas":
				total += item.Values.NaturalGas
			case "productKl":
				total += item.Values.ProductKl
			}
		}
	}
	return total
}
