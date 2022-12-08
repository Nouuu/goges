package kordis

import (
	"github.com/go-resty/resty/v2"
	"github.com/nouuu/goges/conf"
	"os"
	"strings"
	"time"
)

type MygesApi struct {
	username             string
	password             string
	token                string
	tokenType            string
	LastUpdatedTokenDate time.Time
	client               *resty.Client
}

const kordisBaseUrl string = "https://api.kordis.fr"
const kordisConnectUrl = "https://authentication.kordis.fr/oauth/authorize?response_type=token&client_id=skolae-app"
const kordiasAgendaUrl = kordisBaseUrl + "/me/agenda"

func GetMygesCredentials() MygesCredentials {
	return MygesCredentials{
		username: os.Getenv(conf.UsernameEnv),
		password: os.Getenv(conf.PasswordEnv),
	}
	return request.Get(url)
}

func (mygesApi *MygesApi) prepareRequest() *resty.Request {
	r := mygesApi.client.R()
	r.SetHeader("Authorization", mygesApi.tokenType+" "+mygesApi.token)
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Content-Type", "application/json")
	return r
}
