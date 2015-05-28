package main

import (
	"log"
	"time"
)

type NetscalerWatcher struct {
	PollInterval time.Duration
	netscalers   []Netscaler
	pub          chan *NitroResponse
	store        map[string]*NitroResponse
	stop         chan struct{}
}

func CreateWatcher(config *Config) NetscalerWatcher {
	return NetscalerWatcher{
		PollInterval: time.Second * config.PollInterval,
		netscalers:   config.Netscalers,
		store:        make(map[string]*NitroResponse),
	}
}

func (nw *NetscalerWatcher) Run() <-chan *NitroResponse {
	nw.stop = make(chan struct{})
	if nw.pub == nil {
		nw.pub = make(chan *NitroResponse)
	}

	poll := time.Tick(nw.PollInterval)

	go func() {
		for {
			select {
			case <-nw.stop:
				return
			case <-poll:
				go nw.getLatestState()
			}
		}
	}()

	return nw.pub
}
func (nw *NetscalerWatcher) Stop() {
	close(nw.stop)
}
func (nw *NetscalerWatcher) getLatestState() {
	for _, ns := range nw.netscalers {
		for _, group := range ns.Groups {
			if info, err := ns.GetLbGroupInfo(group); err == nil {
				nw.publishChanges(info)
			} else {
				log.Printf("FAILED [%v]", err)
			}
		}
	}
}
func (nw *NetscalerWatcher) publishChanges(current *NitroResponse) {
	if prev, ok := nw.store[current.ServiceName]; ok {
		if prev.State == current.State {
			return
		}
	}
	nw.pub <- current
	nw.store[current.ServiceName] = current
}
