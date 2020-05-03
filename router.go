package main

import (
	"github.com/gorilla/mux"
)

func NewRouter(router *mux.Router) {
	router.Methods("GET").Path("/gh_user").HandlerFunc(GithubHandler)
	router.Methods("GET").Path("/so_badges").HandlerFunc(StackoverflowHandler)
	router.Methods("GET").Path("/").HandlerFunc(IndexHandler)
}
