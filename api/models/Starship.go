package models

// Starship entity
type Starship struct {
	Name                 string   `json:"name"`
	Model                string   `json:"model"`
	Manufacturer         string   `json:"manufacturer"`
	CostInCredits        int      `json:"cost_in_credits"`
	Length               int      `json:"length"`
	MaxAtmospheringSpeed string   `json:"max_atmosphering_speed"`
	Crew                 int      `json:"crew"`
	Passengers           int      `json:"passengers"`
	CargoCapacity        string   `json:"cargo_capacity"`
	Consumables          string   `json:"consumables"`
	HyperdriveRating     float32  `json:"hyperdrive_rating"`
	MGLT                 int      `json:"mglt"`
	StarshipClass        string   `json:"starship_class"`
	Pilots               []People `gorm:"many2many:people_starships; foreignkey:ID" json:"pilots"`
	Films                []Film   `gorm:"many2many:films_starships; foreignkey:ID" json:"films"`
	Created              string   `json:"created"`
	Edited               string   `json:"edited"`
	URL                  int      `json:"url"`
	ID                   int      `gorm:"primary_key"`
}
