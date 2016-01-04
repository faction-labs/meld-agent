package main

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
)

var versionCmd = cli.Command{
	Name:   "agent-version",
	Usage:  "show agent version",
	Action: versionCmdAction,
}

func versionCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	var ver int

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Call("Meld.Version", 0, &ver); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Meld Version: %d\n", ver)
}
