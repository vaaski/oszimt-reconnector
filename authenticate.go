package main

import (
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var jar, _ = cookiejar.New(nil)

func login(username string, password string) {
	log.Println("trying to log in")

	err, loginPage := fetch(LOGIN_ADDR)
	if wouldPanic(err) {
		return
	}

	defer loginPage.Body.Close()

	doc, err := goquery.NewDocumentFromReader(loginPage.Body)
	if wouldPanic(err) {
		return
	}

	// check if already logged in
	logoffButton := doc.Find("[href='http://logoff.now']")
	if logoffButton.Length() > 0 {
		log.Println("already logged in")
		return
	}

	// get token to log in with
	token, exists := doc.Find("input[name=ta_id]").Attr("value")
	if !exists {
		log.Fatalln("no token found")
		return
	}
	log.Println("token:", token)

	cookies := loginPage.Cookies()
	parsedUrl, _ := url.Parse(LOGIN_ADDR)
	jar.SetCookies(parsedUrl, cookies)

	postData := url.Values{}
	postData.Set("ta_id", token)
	postData.Set("uid", username)
	postData.Set("pwd", password)
	postData.Set("voucher_logon_btn", "++Login++")
	encodedData := strings.NewReader(postData.Encode())

	client := &http.Client{Jar: jar}
	loginResponse, loginErr := client.Post(LOGIN_ADDR, "application/x-www-form-urlencoded", encodedData)
	if wouldPanic(loginErr) {
		return
	}

	loginResponseParsed, loginParseError := goquery.NewDocumentFromReader(loginResponse.Body)
	if wouldPanic(loginParseError) {
		return
	}

	errorElement := loginResponseParsed.Find(".message-wrapper.error")
	if errorElement.Length() > 0 {
		log.Println("login error", errorElement.Text())
	}

	log.Println("logged in")
}
