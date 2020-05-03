package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Kaizen API!")
}

func GithubHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.github.com/users/" + r.Header["Username"][0]
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Request failed with error %s\n", err)
	} else {
		buf := new(bytes.Buffer)

		buf.ReadFrom(response.Body)
		newStr := buf.String()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(newStr))
	}
}

func StackoverflowHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://api.stackexchange.com/2.2/users/" + r.Header["User_id"][0] + "/badges?order=desc&sort=rank&site=stackoverflow"
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Request failed with error %s\n", err)
	} else {
		buf := new(bytes.Buffer)

		buf.ReadFrom(response.Body)
		newStr := buf.String()
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(newStr))
	}
}
