package gsonic

import (
	"net"
	"sync"
)

type Ingester struct {
	mu sync.RWMutex
	c  *channel
}

func newIngester(c *Config) (*Ingester, error) {
	i := &Ingester{
		mu: sync.RWMutex{},
	}
	conn, err := newConnection(c, net.Dialer{})
	if err != nil {
		return nil, err
	}
	i.c = conn
	if err = i.c.connect(Ingest); err != nil {
		return nil, err
	}
	return i, nil
}

func (i *Ingester) Push(item *Object) error {
	return i.c.Write(item.Prepare("PUSH"))
}

func (i *Ingester) Pop(item *Object) error {
	return i.c.Write(item.Prepare("POP"))
}

func (i *Ingester) Count() error {
	return nil
}

func (i *Ingester) FlushC() error {
	return nil
}

func (i *Ingester) FlushB() error {
	return nil
}

func (i *Ingester) FlushO() error {
	return nil
}

func (i *Ingester) WriteBulk() <-chan error {
	return nil
}
