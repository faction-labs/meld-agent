package main

import (
	"io/ioutil"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld-agent/agent"
	"github.com/factionlabs/meld-agent/config"
)

var cmdRun = cli.Command{
	Name:   "run",
	Action: runAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Usage: "Path to config file",
			Value: "config.toml",
		},
	},
}

func runAction(c *cli.Context) {
	configPath := c.String("config")

	var data string

	d, err := ioutil.ReadFile(configPath)
	switch {
	case os.IsNotExist(err):
		log.Debug("no config detected; generating local config")
		data = `managerURL = "https://127.0.0.1:8080"
listenAddr = ":9000"
tlsCACert = ""
tlsCert = ""
tlsKey = ""
`
	case err == nil:
		data = string(d)
	default:
		log.Fatal(err)
	}

	f, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}

	f.Write([]byte(data))
	f.Close()

	cfg, err := config.ParseConfig(string(data))
	if err != nil {
		log.Fatal(err)
	}

	agtCfg := &agent.AgentConfig{
		ListenAddr: cfg.ListenAddr,
	}

	agt, err := agent.NewAgent(agtCfg)
	if err != nil {
		log.Fatal(err)
	}

	if err := agt.Run(); err != nil {
		log.Fatal(err)
	}

	waitForInterrupt()
}
