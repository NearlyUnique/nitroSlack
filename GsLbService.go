package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	NitroResponse struct {
		Datacentre  string `json:"datacentre"`
		NitroServer string `json:"nitroServer"`
		ServiceName string `json:"serviceName"`
		State       string `json:"state"`
	}
	gsLbServiceResponse struct {
		Errorcode   int    `json:"errorcode"`
		Message     string `json:"message"`
		Severity    string `json:"severity"`
		GsLbService []struct {
			ServiceName string `json:"serviceName"`
			State       string `json:"state"`
		} `json:"gslbservice"`
	}
	NitroConfig struct {
		Host     string `json:"host"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (nitro NitroConfig) GetLbGroupInfo(group string) (*NitroResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/nitro/v1/stat/gslbservice/%s", nitro.Host, group)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Error : %s", err)
		return nil, err
	}
	req.SetBasicAuth(nitro.Username, nitro.Password)
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error : %s", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error : %s", err)
		return nil, err
	}
	var detail gsLbServiceResponse
	err = json.Unmarshal(body, &detail)

	//log.Printf("RAW: %s\n", string(body))

	if len(detail.GsLbService) == 1 {
		return &NitroResponse{
			NitroServer: nitro.Host,
			ServiceName: detail.GsLbService[0].ServiceName,
			State:       detail.GsLbService[0].State,
		}, err
	}
	return &NitroResponse{}, fmt.Errorf("No content")
}