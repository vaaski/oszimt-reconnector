package main

import (
	"bufio"
	b64 "encoding/base64"
	"fmt"
	"os"
	"strings"
	"syscall"

	"golang.org/x/term"
)

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
