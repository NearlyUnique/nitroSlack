package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var slack_request string

func Test_postToSlack(t *testing.T) {
	ts := httptest.NewServer(respondOk(t))
	cfg := createConfig("http://netscaler.example.com", "xxx", "xxx", ts.URL)

	err := cfg.Slack.PostSlack("Hello")

	AssertNoError(t, err)
	AssertEqualStrings(t,
		`{"text":"Hello","username":"A bot","icon_url":"http://image.example.com/image.png","channel":"#random"}`,
		slack_request)
}

func respondOk(t *testing.T) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//collect request
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			slack_request = string(body)
		} else {
			t.Error("Failed to read slack request")
		}
		w.WriteHeader(http.StatusOK)
	})
}
