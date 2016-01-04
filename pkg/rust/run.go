package rust

import (
	log "github.com/Sirupsen/logrus"
)

// RunCommand runs the specified RCON command and returns the response
func (r *RustServer) RunCommand(command string) (string, error) {
	log.Debugf("rcon.command: cmd=%s", command)
	requestId, err := r.request(command)
	if err != nil {
		return "", err
	}

	// wait for response
	return r.getResponse(requestId)
}
