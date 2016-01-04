package rust

import (
	"fmt"
	"os/exec"

	log "github.com/Sirupsen/logrus"
)

// Start runs a new Rust Server
func (r *RustServer) Start(args *StartArgs) error {
	log.Debugf("starting rust server: args=%v", args)

	pArgs := []string{
		"-batchmode",
		"-nographics",
		"+server.globalchat",
		"true",
		"+server.ip",
		args.ServerIP,
		"+rcon.ip",
		args.RconIP,
		"+server.port",
		fmt.Sprintf("%d", args.ServerPort),
		"+rcon.port",
		fmt.Sprintf("%d", args.RconPort),
		"+rcon.password",
		args.RconPassword,
		"+server.maxplayers",
		fmt.Sprintf("%d", args.MaxPlayers),
		"+server.hostname",
		args.Hostname,
		"+server.identity",
		args.Identity,
		"+server.level",
		args.Level,
		"+server.seed",
		fmt.Sprintf("%d", args.Seed),
		"+server.worldsize",
		fmt.Sprintf("%d", args.WorldSize),
		"+server.saveinterval",
		fmt.Sprintf("%d", args.SaveInterval),
		"+server.description",
		args.Description,
		"+server.url",
		args.URL,
	}

	binPath := RustDedicatedPath()

	c := exec.Command(binPath, pArgs...)

	log.Debugf("starting rust server: cmd=%s args=%v", c.Path, c.Args)

	return nil
}
