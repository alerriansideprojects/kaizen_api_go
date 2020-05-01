package main

import(
	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router){
	router.Methods("GET").Path("/github_user").HandlerFunc(FetchGHUser)
	router.Methods("GET").Path("/").HandlerFunc(FetchIndex)
}
