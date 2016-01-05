package rust

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	ps "github.com/mitchellh/go-ps"
)

func (r *RustServer) Stop() error {
	log.Debug("stopping rust")

	processes, err := ps.Processes()
	if err != nil {
		return err
	}

	for _, proc := range processes {
		if strings.Index(RustDedicatedPath(), proc.Executable()) > -1 {
			log.Debugf("killing rust server: pid=%d", proc.Pid())

			p, err := os.FindProcess(proc.Pid())
			if err != nil {
				return err
			}

			if err := p.Kill(); err != nil {
				return err
			}
		}
	}

	log.Debug("rust server stopped successfully")

	return nil
}
