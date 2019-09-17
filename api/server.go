package api

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // test
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Run test
func Run(addr string) {

	// var film models.Film

	// db.Preloads("Characters").Preloads("Vehicles").Preload("Planets").First(&film)

	// fmt.Printf("film %+v\n", film)

	// var planet models.Planet

	// db.Preloads("Films").First(&planet)

	// fmt.Printf("planet %+v\n", planet)

	// var people models.People

	// db.Preloads("Films").Preload("Starships").Preload("Species").Preload("Vehicles").Find(&people)

	// fmt.Printf("people %+v\n", people)

	// var film models.Film
	// var peoples []models.People

	// db.Model(&film).Related(&peoples, "Characters")
	// db.Preload("Characters").First(&film)

	// db.First(&film)
	// db.Model(&film).Association("Characters").Find(&film.Characters)

	// fmt.Printf("film %+v\n", film)
	//fmt.Printf("peoples %+v\n", peoples)

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

	// db.Close()

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	// r.HandleFunc("/test/{key}", ProductsHandler)

	// r.Use(loggingMiddleware)
	// http.Handle("/", r)

	// http.ListenAndServe(":8080", r)

	// Start http server
	var err error
	var server = Server{}

	server.DB, err := gorm.Open("sqlite3", "swapi.dat")
	if err != nil {
		fmt.Printf("Cannot connect to database")
		log.Fatal("This is the error:", err)
	}

	server.Router = mux.NewRouter()

	//server.initializeRoutes()

	//server.Run(":8080")
}
