package kordis

import (
	"github.com/go-resty/resty/v2"
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
const kordisAgendaUrl = kordisBaseUrl + "/me/agenda"

func (mygesApi *MygesApi) prepareRequest() *resty.Request {
	r := mygesApi.client.R()
	r.SetHeader("Authorization", mygesApi.tokenType+" "+mygesApi.token)
	r.SetHeader("Accept", "application/json")
	r.SetHeader("Content-Type", "application/json")
	return r
}

func (mygesApi *MygesApi) Get(url string, queryParams map[string]string) (*resty.Response, error) {
	request := mygesApi.prepareRequest()
	if queryParams != nil {
		request.SetQueryParams(queryParams)
	}
	return request.Get(url)
}
