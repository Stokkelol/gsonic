package gsonic

import (
	"github.com/rs/zerolog"
	"os"
)

type Config struct {
	Host     string
	Port     string
	Password string
}

type Client struct {
	i   *Ingester
	s   *Searcher
	log zerolog.Logger
}

func NewClient(config *Config, opts []Option) (*Client, error) {
	i, err := newIngester(config)
	if err != nil {
		return nil, err
	}
	s, err := newSearcher(config)
	if err != nil {
		return nil, err
	}

	c := &Client{i: i, s: s, log: zerolog.New(os.Stdout)}
	for i := range opts {
		opts[i](c)
	}
	return c, nil
}

func (c *Client) Searcher() *Searcher {
	return c.s
}

func (c *Client) Ingester() *Ingester {
	return c.i
}
