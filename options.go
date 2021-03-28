package gsonic

import "time"

type Option func(*Client)

func WithDeadline(v time.Time) Option {
	return func(c *Client) {

	}
}

func WithTimeout(v time.Time) Option {
	return func(c *Client) {

	}
}
