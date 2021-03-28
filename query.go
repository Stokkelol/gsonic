package gsonic

import (
	"bytes"
	"fmt"
)

type Query struct {
	Collection string `json:"collection,omitempty"`
	Bucket     string `json:"bucket,omitempty"`
	Term       string `json:"term,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	Lang       Lang   `json:"lang,omitempty"`
}

func (q *Query) Stmt() []byte {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "QUERY %s %s \"%s\"", q.Collection, q.Bucket, q.Term)
	if q.Limit != 0 {
		fmt.Fprintf(b, " LIMIT(%d)", q.Limit)
	}
	if q.Offset != 0 {
		fmt.Fprintf(b, " OFFSET(%d)", q.Offset)
	}
	if q.Lang != "" {
		fmt.Fprintf(b, " LANG(%s)", q.Lang)
	}

	println(b.String())

	return b.Bytes()
}

type Suggestion struct {
	Collection string `json:"collection,omitempty"`
	Bucket     string `json:"bucket,omitempty"`
	Word       string `json:"word,omitempty"`
	Limit      int    `json:"limit,omitempty"`
}

func (s *Suggestion) Stmt() []byte {
	b := new(bytes.Buffer)
	fmt.Fprintf(b, "SUGGEST %s %s \"%s\"", s.Collection, s.Bucket, s.Word)
	if s.Limit != 0 {
		fmt.Fprintf(b, " LIMIT(%d)", s.Limit)
	}
	return b.Bytes()
}
