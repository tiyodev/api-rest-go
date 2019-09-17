package models

// Vehicle entity
type Vehicle struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        int      `gorm:"column:cost_in_credits" json:"cost_in_credits"`
	Length               float32  `json:"length"`
	MaxAtmospheringSpeed int      `gorm:"column:max_atmosphering_speed" json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           int      `json:"passengers"`
	CargoCapacity        int      `gorm:"column:cargo_capacity" json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	VehicleClass         string   `gorm:"column:vehicle_class" json:"vehicle_class"`
	Pilots               []People `json:"pilots"`
	Films                []Film   `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   uint     `gorm:"primary_key" json:"id"`
}
