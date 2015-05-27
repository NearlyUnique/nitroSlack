package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func createConfig(url, group, datacentre string) Config {
	return Config{
		Netscalers: []NetscalerConfig{
			NetscalerConfig{
				NitroConfig: NitroConfig{
					Host:     url,
					Username: "hello",
					Password: "world",
				},
				Datacentre: datacentre,
				Groups:     []string{group},
			},
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
