package meld

import (
	"sync"
)

const (
	Version = 0
)

type Meld struct {
	mu *sync.Mutex
}

func NewMeld() *Meld {
	return &Meld{
		mu: &sync.Mutex{},
	}
}
