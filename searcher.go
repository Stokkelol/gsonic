package gsonic

import (
	"net"
	"strings"
	"sync"
)

const query = "QUERY"
const suggest = "SUGGEST"

type searchable interface {
	Stmt() []byte
}

type Searcher struct {
	mu sync.RWMutex
	c  *channel
}

func newSearcher(c *Config) (*Searcher, error) {
	s := &Searcher{
		mu: sync.RWMutex{},
	}
	conn, err := newConnection(c, net.Dialer{})
	if err != nil {
		return nil, err
	}
	s.c = conn
	if err = s.c.connect(Search); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Searcher) Query(q *Query) ([]string, error) {
	return s.search(q, query)
}

func (s *Searcher) Suggest(sg *Suggestion) ([]string, error) {
	return s.search(sg, suggest)

}

func (s *Searcher) search(sc searchable, typ string) ([]string, error) {
	if err := s.c.Write(sc.Stmt()); err != nil {
		return nil, err
	}

	// STARTED
	resp, err := s.c.Read()
	if err != nil {
		return nil, err
	}

	// PENDING + id of transaction
	resp, err = s.c.Read()
	if err != nil {
		return nil, err
	}

	// EVENT + type
	resp, err = s.c.Read()
	if err != nil {
		return nil, err
	}

	return s.parseResults(resp, typ), nil
}

func (s *Searcher) parseResults(result string, eventType string) []string {
	if strings.HasPrefix(result, s.event(eventType)) {
		return strings.Split(result, " ")[3:]
	}

	return []string{}
}

func (s *Searcher) event(eventType string) string {
	return "EVENT " + eventType
}
