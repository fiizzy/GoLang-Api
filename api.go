package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Creating a Struct for test response
type Test struct {
	Data string
}

// Creating a slice
var dbMocker []string

// dbMocker :=

func main() {
	//define a router
	router := mux.NewRouter()

	router.HandleFunc("/test", test)

	router.HandleFunc("/add/{item}", addItem).Methods("POST")

	//startup the server
	http.ListenAndServe(":5000", router)

}

//This function handles all requests that comes to the /test route
func test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dataStruct := Test{"This is a non Anonymous Struct"}

	json.NewEncoder(w).Encode(dataStruct)

}

func addItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	routeVar := mux.Vars(r)["item"]

	dbMocker = append(dbMocker, routeVar)

	json.NewEncoder(w).Encode(dbMocker)

}
