package rust

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/james4k/rcon"
)

const (
	rustDedicatedWin   = "RustDedicated.exe"
	rustDedicatedLinux = "rust_server"
)

type RustServerConfig struct {
	RconAddress  string
	RconPassword string
}

type RustServer struct {
	config   *RustServerConfig
	console  *rcon.RemoteConsole
	requests map[int]string
	respChan chan int
	errChan  chan error
}

func NewRustServer(config *RustServerConfig) (*RustServer, error) {
	srv := &RustServer{
		config: config,
	}

	// allow nil config to be passed; i.e. to enable starting
	// the initial rust server
	if config != nil {
		c, err := rcon.Dial(config.RconAddress, config.RconPassword)
		if err != nil {
			return nil, fmt.Errorf("error connecting to rcon: %s", err)
		}

		respChan := make(chan int)
		errChan := make(chan error)

		// error reporter channel
		go func() {
			for err := range errChan {
				log.Error(err)
			}
		}()

		srv.console = c
		srv.requests = map[int]string{}
		srv.respChan = respChan
		srv.errChan = errChan

		ticker := time.NewTicker(time.Millisecond * 250)
		// read channel for rcon
		go func() {
			// start read loop
			for _ = range ticker.C {
				resp, id, err := srv.console.Read()
				if err != nil {
					errChan <- err
				}

				srv.requests[id] = resp
				respChan <- id
			}

		}()
	}

	return srv, nil
}
