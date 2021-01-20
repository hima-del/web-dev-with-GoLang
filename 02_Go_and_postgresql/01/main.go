package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type driver struct {
	gorm.Model
	name     string
	liscence string
	cars     []car
}

type car struct {
	gorm.Model
	year      int
	make      string
	modelname string
	driverid  int
}

var db *gorm.DB

var err error

var (
	drivers = []driver{
		{name: "jimmy johnson", liscence: "abc123"},
		{name: "john doe", liscence: "xyz987"},
		{name: "george david", liscence: "etc345"},
	}
	cars = []car{
		{year: 2000, make: "Toyota", modelname: "Tundra", driverid: 1},
		{year: 2001, make: "Honda", modelname: "Accord", driverid: 1},
		{year: 2002, make: "Nissan", modelname: "Sentra", driverid: 2},
		{year: 2003, make: "Ford", modelname: "F-150", driverid: 3},
	}
)

func main() {
	router := mux.NewRouter()
	db, err = gorm.Open("postgres", "host=db port=5432 user=himaja dbname=search sslmode=disable password=password")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.close()
	db.AutoMirate(&driver{})
	db.AutoMirate(&car{})

	for index := range cars {
		db.create(&cars[index])
	}
	for index := range drivers {
		db.create(&drivers[index])
	}

	router.HandleFunc("/cars", getcars).Methods("GET")
	//router.HandleFunc("/cars/{id}", getcar).Methods("GET")
	//router.HandleFunc("/drivers/{id}", getdriver).Methods("GET")
	//router.HandleFunc("/cars/{id}", deletecar).Methods("DELETE")

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func getcars(w http.ResponseWriter, req *http.Request) {
	var cars []car
	db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}
