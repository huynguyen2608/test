package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	// "github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("My website")

	http.HandleFunc("/", ApiTest)
	http.HandleFunc("/about", ApiExam)

	log.Fatal(http.ListenAndServe(":3331", nil))
}

func ApiTest(w http.ResponseWriter, r http.Request) {
	fmt.Print(w, "this is best")
}

func ApiExam(w http.ResponseWriter, r http.Request) {
	fmt.Print(w, "this is bad")
}
