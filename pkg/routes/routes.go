package main

import (
	"log"
	"net/http"

	"../controllers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// User routes
	r.HandleFunc("/users", controllers.GetUsers).Methods("GET")

	// Reservation routes
	r.HandleFunc("/reservations", controllers.GetReservations).Methods("GET")

	// Start the server
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
