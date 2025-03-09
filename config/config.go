package config

type Answers struct {
	Transport []float32 `json:"transport"`
	Food      []float32 `json:"food"`
	Energy    []float32 `json:"energy"`
	Waste     []float32 `json:"waste"`
}

var EmisionFactors = struct {
	TransportEmission []float32
	FoodEmission      []float32
	EnergyEmission    []float32
	WasteEmission     []float32
}{
	TransportEmission: []float32{2.31, 0.1, 0.15, 0.11},
	FoodEmission:      []float32{27, 6.9, 3.2, 2},
	EnergyEmission:    []float32{0.1, 0.202, 1.51, 0.2},
	WasteEmission:     []float32{0.45, 0.1, 6, 1.3},
}
