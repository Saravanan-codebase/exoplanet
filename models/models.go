package models

type Exoplanet struct {
	ID          string  `json:"id,omitempty" `
	Name        string  `json:"name" required:"true"`
	Description string  `json:"description" required:"true" `
	Distance    int     `json:"distance" required:"true"`
	Radius      float64 `json:"radius" required:"true"`
	Mass        float64 `json:"mass,omitempty"`
	Type        string  `json:"type" required:"true"`
}

type FuelEstimation struct {
	FuelCost float64 `json:"fuel_cost"`
}
