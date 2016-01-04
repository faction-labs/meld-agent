package agent

import (
	"net"
	"net/http"
	"net/rpc"
	"runtime"

	log "github.com/Sirupsen/logrus"
	"github.com/factionlabs/meld-agent/agent/meld"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func (a *Agent) Run() error {
	log.Infof("starting agent: addr=%s", a.config.ListenAddr)

	m := meld.NewMeld()
	rpc.Register(m)

	rpc.HandleHTTP()

	l, err := net.Listen("tcp", a.config.ListenAddr)

	if err != nil {
		return err
	}

	go http.Serve(l, nil)

	return nil
}
