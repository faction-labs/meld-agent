package rust

import (
	"bufio"
	"bytes"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
)

func (r *RustServer) Status() (*ServerStatus, error) {
	host, err := r.RunCommand("status")
	if err != nil {
		return nil, err
	}

	status := &ServerStatus{}

	// parse response
	b := bytes.NewBufferString(host)
	scanner := bufio.NewScanner(b)
	for scanner.Scan() {
		t := scanner.Text()

		parts := strings.Split(t, ":")
		if len(parts) < 2 {
			break
		}

		val := strings.Split(parts[1], " ")

		if strings.Index(t, "hostname") > -1 {
			h := strings.Join(val, " ")
			status.Hostname = strings.TrimSpace(h)
		}
		if strings.Index(t, "version") > -1 {
			ver := val[1]
			status.Version = strings.TrimSpace(ver)
		}
		if strings.Index(t, "map") > -1 {
			m := strings.Join(val, " ")
			status.Map = strings.TrimSpace(m)
		}
		if strings.Index(t, "players") > -1 {
			m := val[1]
			p, err := strconv.Atoi(strings.TrimSpace(m))
			if err != nil {
				log.Errorf("unable to parse players from status: %s", err)
				continue
			}
			status.Players = p
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	desc, err := r.Description()
	if err != nil {
		return nil, err
	}

	status.Description = desc

	return status, nil
}
