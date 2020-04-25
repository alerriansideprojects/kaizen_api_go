package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Service...")

	router := mux.NewRouter().StrictSlash(true)

	NewRouter(router)

	log.Fatal(http.ListenAndServe(":8080", router))
	fmt.Println("Service Terminated...")
}
