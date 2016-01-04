package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/factionlabs/meld-agent/agent/meld"
)

var installSteamCmd = cli.Command{
	Name:   "install-steam",
	Usage:  "install steam",
	Action: installSteamCmdAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "steampath, p",
			Usage: "steam path",
			Value: "",
		},
	},
}

func installSteamCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	steamPath := c.String("steampath")

	args := &meld.InstallArgs{
		SteamPath: steamPath,
	}

	var success bool

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Call("Meld.InstallSteam", args, &success); err != nil {
		log.Fatal(err)
	}

	log.Infof("Steam install complete: success=%v", success)
}

var installRustCmd = cli.Command{
	Name:   "install-rust",
	Usage:  "install rust",
	Action: installRustCmdAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "steampath, p",
			Usage: "steam path",
			Value: "",
		},
		cli.StringFlag{
			Name:  "rustpath, r",
			Usage: "rust path",
			Value: "",
		},
	},
}

func installRustCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	steamPath := c.String("steampath")
	rustPath := c.String("rustpath")

	args := &meld.InstallArgs{
		SteamPath: steamPath,
		RustPath:  rustPath,
	}

	var success bool

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Call("Meld.InstallRust", args, &success); err != nil {
		log.Fatal(err)
	}

	log.Infof("Rust install complete: success=%v", success)
}

var installOxideCmd = cli.Command{
	Name:   "install-oxide",
	Usage:  "install rust",
	Action: installOxideCmdAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "rustpath, r",
			Usage: "rust path",
			Value: "",
		},
	},
}

func installOxideCmdAction(c *cli.Context) {
	addr := c.GlobalString("addr")

	rustPath := c.String("rustpath")

	args := &meld.InstallArgs{
		RustPath: rustPath,
	}

	var success bool

	client, err := getClient(addr)
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Call("Meld.InstallOxide", args, &success); err != nil {
		log.Fatal(err)
	}

	log.Infof("Oxide install complete: success=%v", success)
}
