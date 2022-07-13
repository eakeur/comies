package query

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrMandatoryParameter = errors.New("a mandatory parameter was not informed to the query")
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
	if !inCase {
		return q
	}

	return q.append(input, values...)
}

func (q Query) OnlyWhere(inCase bool, input string, values ...interface{}) (Query, error) {
	if !inCase {
		return Query{}, ErrMandatoryParameter
	}

	return q.append(input, values...), nil
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

func (q Query) append(input string, values ...interface{}) Query {
	placeholders := make([]interface{}, len(values))
	for i, value := range values {
		q.Args = append(q.Args, value)
		placeholders[i] = len(q.Args)
	}

	q.query += fmt.Sprintf(q.nextConjunction()+input+" ", placeholders...)

	return q
}
