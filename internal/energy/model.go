package energy

type EnergyValues struct {
	Electricity float64 `json:"electricity"`
	IndexEnergy float64 `json:"indexEnergy"`
	NaturalGas  float64 `json:"naturalGas"`
	ProductKl   float64 `json:"productKl"`
}

type EnergyItem struct {
	Month  string       `json:"month"`
	Year   string       `json:"year"`
	Values EnergyValues `json:"values"`
}

type EnergyResponse struct {
	Data   []EnergyItem `json:"data"`
	Status string       `json:"status"`
}

// Domain helper method (lebih clean daripada switch di banyak tempat)
func (e EnergyItem) GetValueByCategory(category string) float64 {
	switch category {
	case "electricity":
		return e.Values.Electricity
	case "naturalGas":
		return e.Values.NaturalGas
	case "productKl":
		return e.Values.ProductKl
	case "indexEnergy":
		return e.Values.IndexEnergy
	default:
		return 0
	}
}
