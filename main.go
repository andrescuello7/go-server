package main

import (
	"log"
	"net/http"

	"github.com/andrescuello7/go-server/routes"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.Routes(router)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}
	log.Println("Servidor ejecutandose en puerto 3000")
	log.Println(server.ListenAndServe())
}
