package types

import (
	"gorm.io/gorm"
)

type Message struct {
	Status int `json:"status"`
	Info string `json:"info"`
}

type RecommendationResponse struct {
	GlobalRecommendation struct {
		Category   string `json:"category"`
		Suggestion string `json:"suggestion"`
	} `json:"global_recommendation"`
	CategoryRecommendations struct {
		Transport []struct {
			Suggestion string `json:"suggestion"`
		} `json:"transport"`
		Food []struct {
			Suggestion string `json:"suggestion"`
		} `json:"food"`
		Energy []struct {
			Suggestion string `json:"suggestion"`
		} `json:"energy"`
		Waste []struct {
			Suggestion string `json:"suggestion"`
		} `json:"waste"`
	} `json:"category_recommendations"`
	Notes any `json:"notes"` 
}

type CarbonResult struct {
	gorm.Model
	Total float32 
	User_id int
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

type DataRecommendation struct {
	Date      string         `json:"date"`
	Energy    map[string]int `json:"energy"`
	Food      map[string]int `json:"food"`
	Transport map[string]int `json:"transport"`
	Waste     map[string]int `json:"waste"`
	Result    float32 `json:"result"`
}

type FullRecommendation struct {
	DataRecommendation
	RecommendationResponse
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

type Transport struct {
	gorm.Model
	CarKM                float32
	PublicKm             float32
	DomesticFlights      float32
	InternationalFlights float32
	Total                float32
	User_id              int
}

type Food struct {
	gorm.Model
	RedMeat    float32
	WhiteMeat  float32
	Dairy      float32
	Vegetarian float32
	Total      float32
	User_id    int
}

type Energy struct {
	gorm.Model
	ApplianceHours float32
	LightBulbs     float32
	GasTanks       float32
	HvacHours      float32
	Total          float32
	User_id        int
}

type Waste struct {
	gorm.Model
	TrashBags      float32
	FoodWaste      float32
	PlasticBottles float32
	PaperPackages  float32
	Total          float32
	User_id        int
}
