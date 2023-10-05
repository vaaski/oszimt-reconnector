package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func getDoc(url string) *goquery.Document {
	res, err := http.Get("http://wlan-login.oszimt.de/logon/cgi/index.cgi")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

func isLoggedIn() bool {
	doc := getDoc("https://wlan-login.oszimt.de/logon/cgi/index.cgi")

	logoffButton := doc.Find("[href='http://logoff.now']")
	return logoffButton.Length() > 0
}

func main() {
	if isLoggedIn() {
		fmt.Println("logged in")
	} else {
		fmt.Println("not logged in")
	}
}
