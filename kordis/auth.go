package kordis

import (
	"encoding/base64"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/nouuu/goges/conf"
	"net/url"
	"strings"
	"time"
)

func GetMygesApi(c *conf.Config) (*KordisApi, error) {
	credential := &KordisApi{}
	err := credential.connect(c)
	return credential, err
}

func (mygesApi *KordisApi) connect(c *conf.Config) error {

	// Creating a new client and setting the redirect policy to no redirect policy.
	client := resty.New()
	client.SetRedirectPolicy(resty.NoRedirectPolicy())

	// Calling the function `encodedCredentials` on the struct `gesCredentials` this will return an encoded string of the credentials.
	credentials := mygesApi.encodedCredentials(c.Username, c.Password)

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
	mygesApi.token = token
	mygesApi.tokenType = location.Get("token_type")
	mygesApi.LastUpdatedTokenDate = time.Now()

	// This is setting the header for the client.
	mygesApi.client = resty.New()

	return nil
}

func (mygesApi *KordisApi) encodedCredentials(username string, password string) string {
	joined := strings.Join([]string{username, password}, ":")
	bytes := []byte(joined)
	return base64.StdEncoding.EncodeToString(bytes)
}
