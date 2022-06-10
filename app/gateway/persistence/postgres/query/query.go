package query

import (
	"fmt"
	"strings"
)

func NewQuery(script string) Query {
	return Query{
		placeholder: script,
		Args:        []interface{}{},
	}
}

type (
	Query struct {
		Args        []interface{}
		query       string
		placeholder string
		conjunction string
	}
)

func (q Query) Script() string {
	var clause = "where "
	if q.query == "" {
		clause = ""
	}

	script := strings.Trim(
		strings.Replace(
			strings.Replace(q.placeholder, "%query%", q.query, 1),
			"%where_query%", clause+q.query, 1,
		),
		" ",
	)
	return script
}

func (q Query) Where(inCase bool, input string, values ...interface{}) Query {
	if inCase {
		placeholders := make([]interface{}, len(values))
		for i, value := range values {
			q.Args = append(q.Args, value)
			placeholders[i] = len(q.Args)
		}

		q.query += fmt.Sprintf(q.nextConjunction()+input+" ", placeholders...)
	}

	return q
}

func (q Query) And() Query {
	if q.query != "" {
		q.conjunction = "and "
	}

	return q
}

func (q Query) Or() Query {
	if q.query != "" {
		q.conjunction = "or "
	}

	return q
}

func (q Query) nextConjunction() string {
	op := q.conjunction
	return op
}
