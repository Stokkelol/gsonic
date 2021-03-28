package gsonic

import (
	"net"
	"sync"
)

type Controller struct {
	mu sync.RWMutex
	c  *channel
}

func NewController(config *Config) (*Controller, error) {
	c, err := newController(config)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func newController(c *Config) (*Controller, error) {
	i := &Controller{
		mu: sync.RWMutex{},
	}
	conn, err := newConnection(c, net.Dialer{})
	if err != nil {
		return nil, err
	}
	i.c = conn
	if err = i.c.connect(Control); err != nil {
		return nil, err
	}

	return i, nil
}
