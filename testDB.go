package main

import (
	// "encoding/json"
	"fmt"
	// "strconv"
	"os"
	// "strings"
	// "github.com/boltdb/bolt"
	// "github.com/peterhellberg/swapi"
	"github.com/SWMySql/database"
)

func main () {
	dir, _ := os.Getwd()
	fmt.Println("Creat database")
	db, _ := database.StartDB ("SW")
	database.GetInfor(dir + "SW")

	//test func HasObj
	name := "Greedo"
	flag := db.HasObj("people", name)
	fmt.Println(flag)

	//test func SearchByID
	ID := "1"
	fmt.Printf("%s\n", db.SearchByID("people", ID))

	//test func SearchFilmByName
	filmName := "A New Hope"
	fmt.Printf("%s\n", db.SearchFilmByName(filmName))

	//test func SearchPersonByName
	personName := "Luke Skywalker"
	fmt.Printf("%s\n", db.SearchPersonByName(personName))

	//test func SearchPlanetByName
	planetName := "climate"
	fmt.Printf("%s\n", db.SearchPlanetByName(planetName))

	//test func SearchSpeciesByName
	speciesName := "Human"
	fmt.Printf("%s\n", db.SearchSpeciesByName(speciesName))

	//test func SearchStarshipByName
	starshipName := "model"
	fmt.Printf("%s\n", db.SearchStarshipByName(starshipName))

	//test func SearchVehicleByName
	vehicleName := "model"
	fmt.Printf("%s\n", db.SearchVehicleByName([vehicleName))

	//test func SearchByPage
	fmt.Printf("%s\n", db.SearchByPage("people", 1))
}