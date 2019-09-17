package models

// Vehicle entity
type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        string   `json:"cost_in_credits"`
	Length               float32  `json:"length"`
	MaxAtmospheringSpeed int      `json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           int      `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `json:"vehicle_class"`
	Pilots               []People `json:"pilots"`
	Films                []Film   `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   uint     `gorm:"primary_key"`
}
