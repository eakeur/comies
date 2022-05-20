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
		Args         []interface{}
		query        string
		placeholder  string
		nextOperator string
	}
)

func (q Query) Script() string {
	script := fmt.Sprintf(q.placeholder, q.query)

	return strings.Trim(script, " ")
}

func (q Query) Append(inCase bool, input string, values ...interface{}) Query {
	if inCase {
		placeholders := make([]interface{}, len(values))
		for i, value := range values {
			q.Args = append(q.Args, value)
			placeholders[i] = len(q.Args)
		}

		q.getNextOperator()

		q.query += fmt.Sprintf(q.getNextOperator()+input+" ", placeholders...)
	}

	return q
}

func (q Query) And() Query {
	if q.query != "" {
		q.nextOperator = "and "
	}

	return q
}

func (q Query) Or() Query {
	if q.query != "" {
		q.nextOperator = "or "
	}

	return q
}

func (q Query) getNextOperator() string {
	op := q.nextOperator
	q.nextOperator = ""

	return op
}
