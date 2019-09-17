package main

import (
	"fmt"
	"github.com/tiyodev/api-rest-go/api/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// func loggingMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Printf("Middleware de test => ")
// 		// Do stuff here
// 		log.Println(r.RequestURI + "\n")
// 		// Call the next handler, which can be another middleware in the chain, or the final handler.
// 		next.ServeHTTP(w, r)
// 	})
// }

// HomeHandler test
// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Printf("home handler \n")
// 	fmt.Fprintf(w, "home handler \n")
// }

// ProductsHandler test
// func ProductsHandler(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fmt.Printf("test %v", vars["key"])
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Products: %v \n", vars["key"])
// }

func main() {
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
