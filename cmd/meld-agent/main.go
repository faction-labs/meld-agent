package main

import (
	"os"
	"os/signal"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld-agent/version"
)

func waitForInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	for _ = range c {
		// sig is a ^C, handle it
		os.Exit(0)
	}
}

func main() {
	app := cli.NewApp()
	app.Name = version.FullName()
	app.Version = version.FullVersion()
	app.Author = "faction labs"
	app.Usage = "meld game management agent"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "debug, D",
			Usage: "Enable debug logging",
		},
	}
	app.Commands = []cli.Command{
		cmdRun,
	}
	app.Before = func(c *cli.Context) error {
		if c.Bool("debug") {
			log.SetLevel(log.DebugLevel)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
