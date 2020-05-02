package main

import (
	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router) {
	router.Methods("GET").Path("/gh_user").HandlerFunc(FetchGHUser)
	router.Methods("GET").Path("/so_badges").HandlerFunc(FetchSOUser)
	router.Methods("GET").Path("/").HandlerFunc(FetchIndex)
}
