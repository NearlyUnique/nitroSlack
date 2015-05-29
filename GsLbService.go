package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
)

func (ns Netscaler) GetLbGroupInfo(group string) (*NitroResponse, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/nitro/v1/stat/gslbservice/%s", ns.Host, group)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(ns.Username, ns.Password)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var detail gsLbServiceResponse
	err = json.Unmarshal(body, &detail)

	if len(detail.GsLbService) == 1 {
		return &NitroResponse{
			NitroServer: ns.Host,
			ServiceName: detail.GsLbService[0].ServiceName,
			State:       detail.GsLbService[0].State,
			Datacentre:  ns.Datacentre,
		}, err
	}
	return &NitroResponse{}, fmt.Errorf("No content")
}
