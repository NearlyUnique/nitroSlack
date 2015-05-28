package main

import (
	"log"
	"time"
)

type NetscalerWatcher struct {
	PollInterval time.Duration
	config       *Config
	latest       chan *NitroResponse
	pub          chan *NitroResponse
	store        map[string]*NitroResponse
	stop         chan struct{}
}

func CreateWatcher(config *Config) NetscalerWatcher {
	return NetscalerWatcher{
		PollInterval: time.Second * config.PollInterval,
		config:       config,
		latest:       make(chan *NitroResponse),
		store:        make(map[string]*NitroResponse),
	}
}

func (nw *NetscalerWatcher) Run() <-chan *NitroResponse {
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
func (nw *NetscalerWatcher) Stop() {
	close(nw.stop)
}
func (nw *NetscalerWatcher) updateNetscalerStatus() {
	for _, ns := range nw.config.Netscalers {
		for _, group := range ns.Groups {
			if info, err := ns.GetLbGroupInfo(group); err == nil {
				nw.latest <- info
			} else {
				log.Printf("FAILED [%v]", err)
			}
		}
	}
}
func (nw *NetscalerWatcher) updateLatest(current *NitroResponse) {
	if prev, ok := nw.store[current.ServiceName]; ok {
		if prev.State == current.State {
			return
		}
	}
	nw.pub <- current
	nw.store[current.ServiceName] = current
}
