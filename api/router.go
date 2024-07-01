package api

import (
	"arr-sorted/internal/app/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {

	r := mux.NewRouter()
	r.HandleFunc("/check", handlers.CheckHandler)
	r.HandleFunc("/sorted", handlers.Sorted)

	return r
}
