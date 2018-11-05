package auth

import (
	"bufio"
	"log"
	"os"
	"strings"
)

var envUsername string
var envPassword string

type Auth struct {
	Status bool
}

func init() {
	file, err := os.Open("auth/.env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "=")
		if line[0] == "username" {
			envUsername = line[1]
		}
		if line[0] == "password" {
			envPassword = line[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (a *Auth) Authnticate(username string, password string) *Auth {
	if username == envUsername && password == envPassword {
		a.Status = true
	}
	return a
}
