package kordis

import (
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/nouuu/goges/conf"
	"os"
	"strings"
	"time"
)

type MygesCredentials struct {
	username             string
	password             string
	Token                string
	LastUpdatedTokenDate time.Time
}

func (gesCredentials *MygesCredentials) encodedCredentials() string {
	joined := strings.Join([]string{gesCredentials.username, gesCredentials.password}, ":")
	byted := []byte(joined)
	return base64.StdEncoding.EncodeToString(byted)
}

func GetMygesCredentials() MygesCredentials {
	return MygesCredentials{
		username: os.Getenv(conf.UsernameEnv),
		password: os.Getenv(conf.PasswordEnv),
	}
}

func Connect(gesCredentials *MygesCredentials) error {
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())
	credentials := gesCredentials.encodedCredentials()

	resp, _ := client.R().
		EnableTrace().
		SetHeader("Authorization", "Basic "+credentials).
		Get("https://authentication.kordis.fr/oauth/authorize?response_type=token&client_id=skolae-app")

	headers := resp.Header()
	location := headers.Get("Location")
	fmt.Printf("headers = %v\n----\n", headers)
	fmt.Printf("Location = %v\n----\n", location)
	fmt.Printf("resp = %v", resp)
	return nil
}
