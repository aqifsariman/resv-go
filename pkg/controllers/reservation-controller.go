package controllers

import (
	"fmt"
	"net/http"
)

func GetReservations(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of reservations")
}
