package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yagyagoel1/learning_golang/router"
)

func main() {
	fmt.Println("mongodb api")
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", router.Router()))
	fmt.Println("Listening on port 4000")
}
