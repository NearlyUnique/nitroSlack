package main

import (
	"log"
	"time"
)

type NitroWatcher struct {
	PollInterval time.Duration
	config       *Config
	latest       chan *NitroResponse
	pub          chan *NitroResponse
	store        map[string]*NitroResponse
	stop         chan struct{}
}

func CreateWatcher(config *Config) NitroWatcher {
	return NitroWatcher{
		PollInterval: time.Second * 15,
		config:       config,
		latest:       make(chan *NitroResponse),
		store:        make(map[string]*NitroResponse),
	}
}

func (nw *NitroWatcher) Run() <-chan *NitroResponse {
	nw.stop = make(chan struct{})
	nw.pub = make(chan *NitroResponse)

	poll := time.Tick(nw.PollInterval)

	go func() {
		for {
			select {
			case <-nw.stop:
				return
			case <-poll:
				go nw.updateNetscalerStatus()
			case current := <-nw.latest:
				nw.updateLatest(current)
			}
		}
	}()

	return nw.pub
}
func (nw *NitroWatcher) Stop() {
	close(nw.stop)
}
func (nw *NitroWatcher) updateNetscalerStatus() {
	for _, ns := range nw.config.Netscalers {
		for _, group := range ns.Groups {
			if info, err := ns.GetLbGroupInfo(group); err == nil {
				info.Datacentre = ns.Datacentre
				nw.latest <- info
			} else {
				log.Printf("FAILED [%v]", err)
			}
		}
	}
}
func (nw *NitroWatcher) updateLatest(current *NitroResponse) {
	if prev, ok := nw.store[current.ServiceName]; ok {
		if prev.State == current.State {
			return
		}
	}
	nw.pub <- current
	nw.store[current.ServiceName] = current
}
