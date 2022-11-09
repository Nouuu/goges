package kordis

import (
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"goges/conf"
	"net/url"
	"os"
	"strings"
	"time"
)

func (gesCredentials *MygesCredentials) encodedCredentials() string {
	joined := strings.Join([]string{gesCredentials.username, gesCredentials.password}, ":")
	bytes := []byte(joined)
	return base64.StdEncoding.EncodeToString(bytes)
}

func GetMygesCredentials() MygesCredentials {
	return MygesCredentials{
		username: os.Getenv(conf.USERNAME_ENV),
		password: os.Getenv(conf.PASSWORD_ENV),
	}
}

func (gesCredentials *MygesCredentials) Connect() error {
	// Creating a new client and setting the redirect policy to no redirect policy.
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())

	// Calling the function `encodedCredentials` on the struct `gesCredentials` this will return an encoded string of the credentials.
	credentials := gesCredentials.encodedCredentials()

	// This is a request to the kordis api to get a token.
	resp, _ := client.R().
		EnableTrace().
		SetHeader("Authorization", "Basic "+credentials).
		Get(kordisConnectUrl)

	// This is checking if the response status code is not a 3xx.
	if resp.StatusCode()/100 != 3 {
		return fmt.Errorf("error while connecting to kordis: %s", resp.Status())
	}

	// This is parsing the response header to get the token.
	headers := resp.Header()
	location, err := url.ParseQuery(headers.Get("Location"))
	if err != nil {
		return err
	}

	var token string

	for key, value := range location {
		if strings.Contains(key, "access_token") {
			token = value[0]
		}
	}

	if len(token) == 0 {
		return fmt.Errorf("error while parsing token")
	}

	// This is setting the token and the time the token was updated.
	gesCredentials.token = token
	gesCredentials.LastUpdatedTokenDate = time.Now()

	return nil
}
