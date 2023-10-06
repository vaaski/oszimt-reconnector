package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDoc(url string) *goquery.Document {
	res, err := http.Get("http://wlan-login.oszimt.de/logon/cgi/index.cgi")
	maybePanic(err)

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	maybePanic(err)

	return doc
}
