package controllers

import (
	"fmt"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "List of users")
}
