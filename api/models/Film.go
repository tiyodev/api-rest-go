package models

// Film entity
type Film struct {
	Title        string     `json:"title"`
	EpisodeID    int        `json:"episode_id"`
	OpeningCrawl string     `json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	Characters   []People   `gorm:"one2many:films_people; foreignkey:ID" json:"characters"`
	Planets      []Planet   `gorm:"one2many:films_planets; foreignkey:ID" json:"planets"`
	Starships    []Starship `gorm:"one2many:films_starships; foreignkey:ID" json:"starships"`
	Vehicles     []Vehicle  `gorm:"one2many:films_vehicles; foreignkey:ID" json:"vehicles"`
	Species      []Specy    `gorm:"one2many:films_species; foreignkey:ID" json:"species"`
	ReleaseDate  string     `json:"release_date"`
	Dreated      string     `json:"dreated"`
	Edited       string     `json:"edited"`
	URL          int        `json:"url"`
	ID           uint       `gorm:"primary_key"`
}
