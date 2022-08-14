package routes

import (
	"github.com/andrescuello7/go-server/controllers"
	"github.com/gorilla/mux"
)

func Routes(routes *mux.Router) {
	subRoute := routes.PathPrefix("/api").Subrouter()
	subRoute.HandleFunc("/", controllers.GetController).Methods("GET")
}
