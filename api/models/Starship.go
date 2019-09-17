package models

// Starship entity
type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        int      `gorm:"column:cost_in_credits" json:"cost_in_credits"`
	Length               int      `json:"length"`
	MaxAtmospheringSpeed string   `gorm:"column:max_atmosphering_speed" json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           int      `json:"passengers"`
	CargoCapacity        int      `gorm:"column:cargo_capacity" json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     float32  `gorm:"column:hyperdrive_rating" json:"hyperdrive_rating"`
	MGLT                 int      `json:"mglt"`
	StarshipClass        string   `gorm:"column:starship_class" json:"starship_class"`
	Pilots               []People `json:"pilots"`
	Films                []Film   `json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   int      `json:"id"`
}
