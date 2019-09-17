package models

// Film entity
type Film struct {
	Title        string     `json:"title"`
	EpisodeID    int        `gorm:"column:episode_id" json:"episode_id"`
	OpeningCrawl string     `gorm:"column:opening_crawl" json:"opening_crawl"`
	Director     string     `json:"director"`
	Producer     string     `json:"producer"`
	Characters   []People   `gorm:"many2many:films_people;foreignkey:ID" json:"characters"`
	Planets      []Planet   `gorm:"many2many:films_planets;" json:"planets"`
	Starships    []Starship `gorm:"many2many:films_starships;" json:"starships"`
	Vehicles     []Vehicle  `gorm:"many2many:films_vehicles;" json:"vehicles"`
	Species      []Specy    `gorm:"many2many:films_species;" json:"species"`
	ReleaseDate  string     `gorm:"column:release_date" json:"release_date"`
	Dreated      string     `json:"dreated"`
	Edited       string     `json:"edited"`
	URL          int        `json:"url"`
	ID           uint       `gorm:"primary_key" json:"id"`
}
