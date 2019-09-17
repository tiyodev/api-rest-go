package models

// Planet entity
type Planet struct {
	Name           string   `json:"name"`
	RotationPeriod int      `gorm:"column:rotation_period" json:"rotation_period"`
	OrbitalPeriod  int      `gorm:"column:orbital_period" json:"orbital_period"`
	Diameter       int      `json:"Diameter"`
	Climate        string   `json:"climate"`
	Gravity        string   `json:"gravity"`
	Terrain        string   `json:"terrain"`
	SurfaceWater   int      `gorm:"column:surface_water" json:"surface_water"`
	Population     int      `json:"population"`
	Residents      []People `json:"residents"`
	Films          []Film   `json:"films"`
	Created        string   `json:"created"`
	Edited         string   `json:"edited"`
	URL            int      `json:"url"`
	ID             uint     `gorm:"primary_key" json:"id"`
}
