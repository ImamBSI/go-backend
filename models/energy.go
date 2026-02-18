package models

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
