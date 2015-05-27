package main

import (
	"bytes"
	"log"
	"os"
	"os/signal"
	"syscall"
	"text/template"
)

var (
	config *Config
	tmpl   *template.Template
)

type NitroResponseChan chan *NitroResponse

func init() {
	var err error
	if config, err = LoadConfig("config.json"); err != nil {
		log.Fatalf("Error reading config.json\n%v\n", err)
	}
	tmpl, err = template.New("slack").Parse(config.Slack.Template.Text)
	if err != nil {
		log.Fatalf("Failed to parse template\n%v\n", err)
	}
	tmpl.Execute(os.Stderr, config)
}

func main() {
	log.Printf("Starting...")
	var buf bytes.Buffer

	sigs := make(chan os.Signal)
	watcher := CreateWatcher(config)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	ch := watcher.Run()

	for {
		select {
		case <-sigs:
			log.Println("\nCTRL-C exiting ...")
			os.Exit(0)
		case change := <-ch:
			log.Printf("\nTell Slack: %s : %s", change.ServiceName, change.State)

			msg := config.Slack.Template

			tmpl.Execute(&buf, change)
			msg.Text = buf.String()

			log.Printf(">%s", msg.Text)

			msg.PostSlackMessage(config.Slack.Url)
		}
	}
}
