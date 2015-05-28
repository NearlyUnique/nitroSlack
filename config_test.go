package main

import "testing"

func Test_config_can_be_loaded(t *testing.T) {
	cfg, err := LoadConfig("config-sample.json")

	AssertNoError(t, err)

	AssertEqualInts(t, 2, len(cfg.Netscalers))

	AssertEqualStrings(t, "dc1", cfg.Netscalers[0].Datacentre)
	AssertEqualStrings(t, "https://dd1.company.com", cfg.Netscalers[0].Host)
	AssertEqualStrings(t, "readonly1", cfg.Netscalers[0].Username)
	AssertEqualStrings(t, "secret1", cfg.Netscalers[0].Password)
	AssertEqualStrings(t, "group1-1", cfg.Netscalers[0].Groups[0])
	AssertEqualStrings(t, "group1-2", cfg.Netscalers[0].Groups[1])

	AssertEqualStrings(t, "https://slack.example.com/random/number", cfg.Slack.Url)

	AssertEqualStrings(t, "`{{.Datacentre}}` is now *{{.State}}*", cfg.Slack.Template)
	AssertEqualStrings(t, "Sample Name", cfg.Slack.Username)
	AssertEqualStrings(t, "http://golang.org/doc/gopher/doc.png", cfg.Slack.IconUrl)
	AssertEqualStrings(t, "#sample", cfg.Slack.Channel)
}
