package rust

import (
	log "github.com/Sirupsen/logrus"
)

func (r *RustServer) request(cmd string) (int, error) {
	requestId, err := r.console.Write(cmd)
	if err != nil {
		return -1, err
	}

	log.Debugf("rcon request id: %d", requestId)

	// add to requests
	r.requests[requestId] = ""

	return requestId, err
}
