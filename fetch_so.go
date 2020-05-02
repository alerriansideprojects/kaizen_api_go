package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func FetchSOUser(w http.ResponseWriter, r *http.Request) {
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
