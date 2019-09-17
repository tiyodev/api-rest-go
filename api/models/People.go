package models

// People entity
type People struct {
	Name      string     `json:"name"`
	Height    string     `json:"height"`
	Mass      string     `json:"mass"`
	HairColor string     `json:"hair_color"`
	SkinColor string     `json:"skin_color"`
	EyeColor  string     `json:"eye_color"`
	BirthYear string     `json:"birth_year"`
	Gender    string     `json:"gender"`
	Homeworld int        `json:"homeworld"`
	Films     []Film     `gorm:"one2many:films_people; foreignkey:ID" json:"films"`
	Species   []Specy    `gorm:"one2many:people_species; foreignkey:ID" json:"species"`
	Vehicles  []Vehicle  `gorm:"one2many:people_vehicles; foreignkey:ID" json:"vehicles"`
	Starships []Starship `gorm:"one2many:people_starships; foreignkey:ID" json:"starships"`
	Created   string     `json:"created"`
	Edited    string     `json:"edited"`
	URL       int        `json:"url"`
	ID        uint       `gorm:"primary_key"`
}

// TableName Get table name
func (People) TableName() string {
	return "people"
}
