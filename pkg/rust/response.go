package rust

import (
	"fmt"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
)

func (r *RustServer) getResponse(requestId int) (string, error) {
	// wait for response
	resp := ""
	started := time.Now()
	// 1 min timeout
	timeout := started.Add(time.Minute * 1)
	for respId := range r.respChan {
		if respId == requestId {
			resp = r.requests[respId]
			delete(r.requests, respId)
			break
		}

		// TODO: run in goroutine to not wait on channel msg
		if time.Now().After(timeout) {
			log.Warnf("getResponse timeout: id=%d", requestId)
			resp = "timeout"
			delete(r.requests, respId)
			break
		}
	}

	return resp, nil
}

func (r *RustServer) parseStandardResponse(resp string) (string, error) {
	parts := strings.Split(resp, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("unable to determine hostname")
	}

	h := parts[1]

	h = strings.Replace(parts[1], "\"", "", -1)
	h = strings.TrimSpace(h)

	return h, nil
}
