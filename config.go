package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type (
	Config struct {
		PollInterval time.Duration `json:"pollInterval"`
		Netscalers   []Netscaler   `json:"netscalers"`
		Slack        struct {
			Url      string `json:"url"`
			Template string `json:"template"`
			Username string `json:"username"`
			IconUrl  string `json:"iconUrl"`
			Channel  string `json:"channel"`
		}
	}
	Netscaler struct {
		Host       string   `json:"host"`
		Username   string   `json:"username"`
		Password   string   `json:"password"`
		Datacentre string   `json:"datacentre"`
		Groups     []string `json:"groups"`
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
