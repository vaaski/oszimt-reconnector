package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/term"
)

var home, homeErr = os.UserHomeDir()
var CREDENTAIL_FILE = filepath.Join(home, ".oszimt-credentials")

func maybePanic(e error) {
	if e != nil {
		panic(e)
	}
}

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

func isLoggedIn() bool {
	doc := getDoc("https://wlan-login.oszimt.de/logon/cgi/index.cgi")

	logoffButton := doc.Find("[href='http://logoff.now']")
	return logoffButton.Length() > 0
}

func readCredentials() (string, string, error) {
	data, err := os.ReadFile(CREDENTAIL_FILE)
	if err != nil {
		return "", "", err
	}

	bDecoded, decodeErr := b64.StdEncoding.DecodeString(string(data))
	if decodeErr != nil {
		return "", "", decodeErr
	}

	decoded := string(bDecoded)
	credentials := strings.Split(string(decoded), ":")
	return credentials[0], credentials[1], nil
}

func askForCredentials() {
	fmt.Print("username: ")
	reader := bufio.NewReader(os.Stdin)
	bUsername, err := reader.ReadString('\n')
	maybePanic(err)
	username := strings.TrimSpace(string(bUsername))

	fmt.Print("password: ")
	bPassword, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	maybePanic(err)
	password := strings.TrimSpace(string(bPassword))

	joined := strings.Join([]string{username, password}, ":")
	encoded := b64.StdEncoding.EncodeToString([]byte(joined))
	writeErr := os.WriteFile(CREDENTAIL_FILE, []byte(encoded), 0644)
	maybePanic(writeErr)

	fmt.Println("credentials saved at", CREDENTAIL_FILE)
}

func main() {
	// if isLoggedIn() {
	// 	log.Println("logged in")
	// } else {
	// 	log.Println("not logged in")
	// }

	username, password, err := readCredentials()
	if err != nil {
		askForCredentials()
		username, password, err = readCredentials()
		maybePanic(err)
	}

	log.Println(username, password)
}
