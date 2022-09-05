package main

import (
	"context"
	"flag"
	"fmt"
	"strings"

	icloudgo "github.com/alvaroaleman/icloud-go"
)

func main() {
	user := flag.String("user", "", "Apple ID")
	flag.Parse()

	if err := icloudgo.Login(context.Background(), *user, prompt("Password:"), prompt("2FA code:")); err != nil {
		fmt.Printf("login failed: %v\n", err)
	}
}

func prompt(toPrint string) func() (string, error) {
	return func() (string, error) {
		fmt.Println(toPrint)
		var result string
		if _, err := fmt.Scanln(&result); err != nil {
			return "", err
		}

		return strings.TrimSpace(result), nil
	}
}
