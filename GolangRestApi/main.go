package main

import (
	"GolangRestApi/userapi"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/v1/user/find", userapi.FindUser).Methods("GET")
	router.HandleFunc("/api/v1/user/getall", userapi.GetAll).Methods("GET")
	router.HandleFunc("/api/v1/user/create", userapi.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/update", userapi.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/delete", userapi.Delete).Methods("DELETE")

	fmt.Printf("Golang Rest API Is Running On Port: 5000")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		panic(err)
	}

}
