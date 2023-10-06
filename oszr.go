package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
)

var home, homeErr = os.UserHomeDir()
var CREDENTAIL_FILE = filepath.Join(home, ".oszimt-credentials")

const LOGIN_ADDR = "https://wlan-login.oszimt.de/logon/cgi/index.cgi"

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}

func isLoggedIn() bool {
	doc := getDoc(LOGIN_ADDR)

	logoffButton := doc.Find("[href='http://logoff.now']")
	return logoffButton.Length() > 0
}

func main() {
	username, password, err := readCredentials()
	if err != nil {
		askForCredentials()
		username, password, err = readCredentials()
		maybePanic(err)
	}

	r1 := regexp.MustCompile(`.`)
	hiddenPassword := r1.ReplaceAllString(password, "*")
	log.Println(username, hiddenPassword)

	if isLoggedIn() {
		log.Println("logged in")
	} else {
		log.Println("not logged in")
	}
}
