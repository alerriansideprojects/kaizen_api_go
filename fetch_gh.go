package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func FetchGHUser(w http.ResponseWriter, r *http.Request) {
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
