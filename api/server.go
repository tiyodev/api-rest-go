package api

import (
	"fmt"
	"github.com/tiyodev/api-rest-go/api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // test
)

// Run test
func Run() {
	db, err := gorm.Open("sqlite3", "swapi.dat")
	if err != nil {
		panic("failed to connect database")
	}

	var film Film
	var peoples []People

	db.Model(&film).Related(&peoples, "Characters")
	db.Preload("Characters").First(&film)

	// db.First(&film)
	// db.Model(&film).Association("Characters").Find(&film.Characters)

	fmt.Printf("film %+v\n", film)
	fmt.Printf("film %+v\n", peoples)

	db.Close()

	// var films []Film
	// db.Table("films").Find(&films)

	// fmt.Printf("People %+v\n", db.HasTable("films"))

	// var film Film
	// db.Raw("SELECT * FROM films WHERE id = ?", 4).Scan(&film)

	// Read
	// film := Film{}
	// errors := db.Table("films").First(film, 1).GetErrors() // find product with id 1

	// fmt.Printf("Film %+v\n", film)

	// fmt.Println(len(errors))

	// for _, err := range errors {
	// 	fmt.Println(err)
	// }

	defer db.Close()

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/test/{key}", ProductsHandler)

	// r.Use(loggingMiddleware)
	// http.Handle("/", r)

	// http.ListenAndServe(":8080", r)

}