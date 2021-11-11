package main

import (
	"log"
	"mux-api/routes"
	"net/http"

	// "net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	const port string = ":5001"
	
	r.HandleFunc("/gets",routes.GetUsers).Methods("GET")
	r.HandleFunc("/posts", routes.AddUsers).Methods("POST")

	log.Println("Server running...", port)
	log.Fatal(http.ListenAndServe(port,r))
}

