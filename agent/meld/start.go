package meld

import (
	"github.com/factionlabs/meld-agent/pkg/rust"
)

func (m *Meld) StartRust(args *rust.StartArgs, pid *int) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	r, err := rust.NewRustServer(nil)
	if err != nil {
		*pid = -1
		return err
	}

	p, err := r.Start(args)
	if err != nil {
		*pid = -1
		return err
	}

	*pid = p
	return nil
}
