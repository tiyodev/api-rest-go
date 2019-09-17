package models

// People entity
type People struct {
	Name      string     `json:"name"`
	Height    int        `json:"height"`
	Mass      int        `json:"mass"`
	HairColor string     `gorm:"column:hair_color" json:"hair_color"`
	SkinColor string     `gorm:"column:skin_color" json:"skin_color"`
	EyeColor  string     `gorm:"column:eye_color" json:"eye_color"`
	BirthYear string     `gorm:"column:birth_year" json:"birth_year"`
	Gender    string     `json:"gender"`
	Homeworld int        `json:"homeworld"`
	Films     []Film     `gorm:"many2many:films_people;foreignkey:ID" json:"films"`
	Species   []Specy    `json:"species"`
	Vehicles  []Vehicle  `json:"vehicles"`
	Starships []Starship `json:"starships"`
	Created   string     `json:"created"`
	Edited    string     `json:"edited"`
	URL       int        `json:"url"`
	ID        uint       `gorm:"primary_key" json:"id"`
}

// TableName Get table name
func (People) TableName() string {
	return "people"
}
