package main

import (
	"log"
	"net/http"
)

func fetch(url string) *http.Response {
	res, err := http.Get(url)
	maybePanic(err)

	// defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	return res
}
