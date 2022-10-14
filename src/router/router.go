package router

import (
	"github.com/gorilla/mux"

	"github.com/MydroX/stickers-collection/src/handlers"
)

func New() *mux.Router {
	r := mux.NewRouter()

	// User
	r.HandleFunc("/team/create", handlers.CreateTeam).Methods("POST")

	return r
}
