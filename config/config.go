package config

import "time"

type Transport struct {
	R1      float32 `gorm:"type:float"`
	R2      float32 `gorm:"type:float"`
	R3      float32 `gorm:"type:float"`
	R4      float32 `gorm:"type:float"`
	Total   float32 `gorm:"type:float"`
	User_id int
	Date    time.Time
}

type Food struct {
	R1      float32 `gorm:"type:float"`
	R2      float32 `gorm:"type:float"`
	R3      float32 `gorm:"type:float"`
	R4      float32 `gorm:"type:float"`
	Total   float32 `gorm:"type:float"`
	User_id int
	Date    time.Time
}

type Energy struct {
	R1      float32 `gorm:"type:float"`
	R2      float32 `gorm:"type:float"`
	R3      float32 `gorm:"type:float"`
	R4      float32 `gorm:"type:float"`
	Total   float32 `gorm:"type:float"`
	User_id int
	Date    time.Time
}

type Waste struct {
	R1      float32 `gorm:"type:float"`
	R2      float32 `gorm:"type:float"`
	R3      float32 `gorm:"type:float"`
	R4      float32 `gorm:"type:float"`
	Total   float32 `gorm:"type:float"`
	User_id int
	Date    time.Time
}

type Answers struct {
	Transport []float32 `json:"transport"`
	Food      []float32 `json:"food"`
	Energy    []float32 `json:"energy"`
	Waste     []float32 `json:"waste"`
}

type ArrayData struct {
	Array []float32
	Index int
}

type Data struct {
	Date      string         `json:"date"`
	Energy    map[string]int `json:"energy"`
	Food      map[string]int `json:"food"`
	Transport map[string]int `json:"transport"`
	Waste     map[string]int `json:"waste"`
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
