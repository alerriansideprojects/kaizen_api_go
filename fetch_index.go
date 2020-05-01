package main

import (
	"fmt"
	"net/http"
)

func FetchIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
