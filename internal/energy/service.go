package energy

type YearlyTotal struct {
	Year   string      `json:"year"`
	Values EnergyValues `json:"values"`
}

// SumAllEnergyByYear mengembalikan total energi per tahun untuk semua kategori
func (s *Service) SumAllEnergyByYear() []YearlyTotal {
	yearMap := make(map[string]EnergyValues)
	for _, item := range s.data {
		y := item.Year
		v := yearMap[y]
		v.Electricity += item.Values.Electricity
		v.IndexEnergy += item.Values.IndexEnergy
		v.NaturalGas += item.Values.NaturalGas
		v.ProductKl += item.Values.ProductKl
		yearMap[y] = v
	}
	var result []YearlyTotal
	for year, values := range yearMap {
		result = append(result, YearlyTotal{
			Year:   year,
			Values: values,
		})
	}
	// Optional: sort by year ascending
	// sort.Slice(result, func(i, j int) bool { return result[i].Year < result[j].Year })
	return result
}
type Service struct {
	data []EnergyItem
}

func NewService(data []EnergyItem) *Service {
	return &Service{data: data}
}

type FilteredEnergy struct {
	Month    string  `json:"month"`
	Year     string  `json:"year"`
	Category string  `json:"category"`
	Value    float64 `json:"value"`
}

// GetRawEnergy mengembalikan data sesuai filter
func (s *Service) GetRawEnergy(year string, category string) interface{} {
	var filtered []EnergyItem

	// Filter by year
	if year != "" {
		for _, item := range s.data {
			if item.Year == year {
				filtered = append(filtered, item)
			}
		}
	} else {
		filtered = s.data
	}

	// Jika tidak ada category → return full data
	if category == "" {
		return filtered
	}

	// Jika ada category → return structured response
	var result []FilteredEnergy
	for _, item := range filtered {
		result = append(result, FilteredEnergy{
			Month:    item.Month,
			Year:     item.Year,
			Category: category,
			Value:    item.GetValueByCategory(category),
		})
	}

	return result
}

// SumEnergy menghitung total
func (s *Service) SumEnergy(year string, category string) float64 {
	var total float64

	for _, item := range s.data {
		if year == "" || item.Year == year {
			total += item.GetValueByCategory(category)
		}
	}

	return total
}
