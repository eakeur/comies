package throw

import (
	"fmt"
)

type (
	fault struct {
		parameters  map[string]interface{}
		description string
		operation   string
		line        int
		inner       error
		child       error
	}

	DetailedError interface {
		Describe(description string) DetailedError
		DescribeF(description string, a ...interface{}) DetailedError
		Params(params map[string]interface{}) DetailedError
		Wrap() DetailedError
		Error() string
		Unwrap() error
	}
)

func (e fault) Describe(description string) DetailedError {
	if e.description == "" {
		e.description = description
	}

	return e
}

func (e fault) DescribeF(description string, a ...interface{}) DetailedError {
	if e.description == "" {
		e.description = fmt.Sprintf(description, a...)
	}

	return e
}

func (e fault) Params(params map[string]interface{}) DetailedError {
	if e.parameters == nil {
		e.parameters = params
	}

	return e
}

func (e fault) Wrap() DetailedError {
	return wrap(e, 2)
}

func (e fault) Error() string {

	var params string
	if e.parameters != nil {
		params = fmt.Sprint(e.parameters)
		params = params[4 : len(params)-1]
	}
	return fmt.Sprintf("%s#%d(%s): %s ----> %v", e.operation, e.line, params, e.description, e.child.Error())
}

func (e fault) Unwrap() error {
	return e.child
}
