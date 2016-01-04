package rust

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld-agent/utils"
)

// Start runs a new Rust Server
func (r *RustServer) Start(args *StartArgs) (int, error) {
	log.Debugf("starting rust server: args=%v", args)

	pArgs := []string{
		"-batchmode",
		//"-nographics",
		"+server.globalchat",
		"true",
		"+server.ip",
		strconv.Quote(args.ServerIP),
		"+rcon.ip",
		strconv.Quote(args.RconIP),
		"+server.port",
		fmt.Sprintf("%d", args.ServerPort),
		"+rcon.port",
		fmt.Sprintf("%d", args.RconPort),
		"+rcon.password",
		strconv.Quote(args.RconPassword),
		"+server.maxplayers",
		fmt.Sprintf("%d", args.MaxPlayers),
		"+server.hostname",
		strconv.Quote(args.Hostname),
		"+server.identity",
		strconv.Quote(args.Identity),
		"+server.level",
		strconv.Quote(args.Level),
		"+server.seed",
		fmt.Sprintf("%d", args.Seed),
		"+server.worldsize",
		fmt.Sprintf("%d", args.WorldSize),
		"+server.saveinterval",
		fmt.Sprintf("%d", args.SaveInterval),
		"+server.description",
		strconv.Quote(args.Description),
		"+server.url",
		strconv.Quote(args.URL),
	}

	binPath := RustDedicatedPath()

	c := exec.Command(binPath, pArgs...)
	c.Dir = utils.GetRustDir()

	log.Debugf("starting rust server: cmd=%s args=%v", c.Path, c.Args)

	if err := c.Start(); err != nil {
		return -1, err
	}

	// wait slightly for process to start
	time.Sleep(time.Millisecond * 500)

	pid := c.Process.Pid

	log.Debugf("rust server started: pid=%d", pid)

	return pid, nil
}
