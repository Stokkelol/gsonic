package gsonic

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
	"unicode"
)

const ERR = "ERR "
const STARTED = "STARTED "
const CONNECTED = "CONNECTED "

const START = "START"
const QUIT = "QUIT"

type channel struct {
	config    *Config
	conn      net.Conn
	reader    *bufio.Reader
	closed    bool
	maxBytes  int
	connected bool
}

func newConnection(c *Config, dialer net.Dialer) (*channel, error) {
	conn, err := dialer.Dial("tcp", fmt.Sprintf("%s:%s", c.Host, c.Port))
	if err != nil {
		return nil, err
	}

	return &channel{
		config: c,
		conn:   conn,
		reader: bufio.NewReader(conn),
		closed: false,
	}, nil
}

func (c *channel) connect(mode Channel) error {
	err := c.Write([]byte(c.prepareConnect(mode)))
	if err != nil {
		return err
	}

	_, err = c.Read()
	return err
}

func (c *channel) Read() (string, error) {
	if c.closed {
		return "", errConnClosed
	}
	buff, err := c.read()
	if err != nil {
		return "", err
	}
	str := buff.String()
	if err = c.parseRead(str); err != nil {
		return "", err
	}
	return str, nil
}

func (c *channel) Write(payload []byte) error {
	payload = append(payload, []byte("\r\n")...)
	_, err := c.conn.Write(payload)
	return err
}

func (c *channel) Close() error {
	return c.close()
}

func (c *channel) read() (*bytes.Buffer, error) {
	buff := &bytes.Buffer{}
	for {
		line, isPrefix, err := c.reader.ReadLine()
		buff.Write(line)
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		if !isPrefix {
			break
		}
	}

	return buff, nil
}

func (c *channel) parseRead(str string) error {
	if strings.HasPrefix(str, ERR) {
		return fmt.Errorf("%s", str[4:])
	}
	if strings.HasPrefix(str, STARTED) {
		ss := strings.FieldsFunc(str, func(r rune) bool {
			if unicode.IsSpace(r) || r == '(' || r == ')' {
				return true
			}
			return false
		})
		bufferSize, err := strconv.Atoi(ss[len(ss)-1])
		if err != nil {
			return errors.New(fmt.Sprintf("Unable to parse STARTED response: %s", str))
		}
		c.maxBytes = bufferSize
	}
	if strings.HasPrefix(str, CONNECTED) {
		c.connected = true
	}

	return nil
}

func (c *channel) prepareConnect(mode Channel) string {
	b := new(strings.Builder)
	fmt.Fprintf(b, "START %s %s", mode, c.config.Password)
	return b.String()
}

func (c *channel) close() error {
	_ = c.Write([]byte("QUIT"))
	if c.closed && c.conn == nil {
		return nil
	}

	c.closed = true
	c.conn = nil
	c.reader = nil
	return nil
}
