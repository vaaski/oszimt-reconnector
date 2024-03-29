//go:generate go-winres make
//go:generate goreleaser --clean --snapshot
//go:generate go run mac-bundle/main.go

package main

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

var home, homeErr = os.UserHomeDir()
var CREDENTAIL_FILE = filepath.Join(home, ".oszimt-credentials")

const LOGIN_ADDR = "http://wlan-login.oszimt.de/logon/cgi/index.cgi"

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}
func wouldPanic(e error) bool {
	if e != nil {
		log.Println(e)
		return true
	}

	return false
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
	log.Println("username", username)
	log.Println("password", hiddenPassword)

	for {
		login(username, password)
		time.Sleep(3 * time.Second)
	}
}
