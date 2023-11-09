package main

import (
	"errors"
	"log"
	"net/http"
)

func fetch(url string) (error, *http.Response) {
	res, err := http.Get(url)
	if err != nil {
		return err, nil
	}

	// defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return errors.New("status code wasn't 200"), nil
	}

	return nil, res
}
