package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	SlackMessage struct {
		Text     string `json:"text"`
		Username string `json:"username"`
		IconUrl  string `json:"icon_url"`
		Channel  string `json:"channel"`
	}
)

func (cfg *Config) CreateMsg(text string) *SlackMessage {
	return &SlackMessage{
		text,
		cfg.Slack.Username,
		cfg.Slack.IconUrl,
		cfg.Slack.Channel,
	}
}
func (s *SlackMessage) PostSlackMessage(url string) error {
	if url == "debug" {
		log.Printf("HTTP POST -> Slack\n%v\n", *s)
	} else {
		jsonStr, _ := json.Marshal(&s)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}

		if resp, err := client.Do(req); err == nil {
			defer resp.Body.Close()

			if body, err := ioutil.ReadAll(resp.Body); err != nil {
				log.Printf("Err:%v\nBody: %s\n", err, body)
				return err
			}

		} else {
			return err
		}
	}
	return nil
}
