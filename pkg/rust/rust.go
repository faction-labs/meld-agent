package rust

import (
	"fmt"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/james4k/rcon"
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
	c, err := rcon.Dial(config.RconAddress, config.RconPassword)
	if err != nil {
		return nil, fmt.Errorf("error connecting to rcon: %s", err)
	}

	respChan := make(chan int)
	errChan := make(chan error)

	srv := &RustServer{
		config:   config,
		console:  c,
		requests: map[int]string{},
		respChan: respChan,
		errChan:  errChan,
	}

	// error reporter channel
	go func() {
		for err := range errChan {
			log.Error(err)
		}
	}()

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

	return srv, nil
}
