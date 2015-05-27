package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type (
	Config struct {
		Netscalers []NetscalerConfig
		Slack      struct {
			Url      string       `json:"url"`
			Template SlackMessage `json:"template"`
		}
	}
	NetscalerConfig struct {
		NitroConfig
		Groups     []string `json:"groups"`
		Datacentre string   `json:"datacentre"`
	}
)

func LoadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("ERROR:%s", err)
	}
	cfg := &Config{}
	if err == nil {
		err = json.Unmarshal(buf, cfg)
	}
	return cfg, err
}
