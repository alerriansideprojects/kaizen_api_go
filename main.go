package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting Service...")

	router := mux.NewRouter().StrictSlash(true)

	NewRouter(router)

	err := http.ListenAndServe(GetPort(), router)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("Service Terminated...")
}

func GetPort() string {
	var port = os.Getenv("PORT")

	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	return ":" + port
}
