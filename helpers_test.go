package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func createConfig(url, group, datacentre, slackUrl string) Config {
	return Config{
		PollInterval: time.Second * 1,
		Netscalers: []Netscaler{
			Netscaler{
				Host:       url,
				Username:   "hello",
				Password:   "world",
				Datacentre: datacentre,
				Groups:     []string{group},
			},
		},
		Slack: SlackConfig{
			Url:      slackUrl,
			Template: "Test",
			Username: "A bot",
			IconUrl:  "http://image.example.com/image.png",
			Channel:  "#random",
		},
	}
}

func writeJson(contentPath string, httpStatus int) http.HandlerFunc {
	var content []byte
	var err error
	if len(contentPath) > 0 {
		if content, err = ioutil.ReadFile(contentPath); err != nil {
			panic(err)
		}
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		if len(content) > 0 {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprint(w, string(content))
		}
	})
}
