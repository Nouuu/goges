package kordis

import (
	"time"
)

type MygesCredentials struct {
	username             string
	password             string
	token                string
	LastUpdatedTokenDate time.Time
}

const kordisBaseUrl string = "https://authentication.kordis.fr"
const kordisConnectUrl = kordisBaseUrl + "/oauth/authorize?response_type=token&client_id=skolae-app"
