package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_basicRequest(t *testing.T) {
	ts := httptest.NewServer(writeJson("dataCentre1-happy.json", http.StatusOK))
	cfg := createConfig(ts.URL, "group1", "data-centre1")

	resp, err := cfg.Netscalers[0].GetLbGroupInfo("data-centre1")

	AssertNoError(t, err)
	AssertEqualStrings(t, "UP", resp.State)
}
