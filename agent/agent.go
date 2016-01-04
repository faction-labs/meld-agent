package agent

import (
	"sync"
)

type Agent struct {
	config *AgentConfig
	mu     *sync.Mutex
}

type AgentConfig struct {
	ListenAddr string
}

func NewAgent(cfg *AgentConfig) (*Agent, error) {
	agt := &Agent{
		config: cfg,
		mu:     &sync.Mutex{},
	}

	return agt, nil
}
