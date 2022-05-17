package fault

import (
	"encoding/json"
	"fmt"
	"strings"
)

type (
	fault struct {
		parameters  map[string]interface{}
		description string
		operation   string
		file        string
		line        int
		inner       error
		child       error
	}

	Error interface {
		Describe(description string) Error
		DescribeF(description string, a ...interface{}) Error
		Params(params map[string]interface{}) Error
		Wrap() Error
		Error() string
		Unwrap() error
	}
)

func (e fault) Describe(description string) Error {
	if e.description == "" {
		e.description = description
	}

	return e
}

func (e fault) DescribeF(description string, a ...interface{}) Error {
	if e.description == "" {
		e.description = fmt.Sprintf(description, a...)
	}

	return e
}

func (e fault) Params(params map[string]interface{}) Error {
	if e.parameters == nil {
		e.parameters = params
	}

	return e
}

func (e fault) Wrap() Error {
	return wrap(e, 2)
}

func (e fault) Error() string {

	params := "{}"
	res, err := json.Marshal(e.parameters)
	if err == nil && string(res) != "null" {
		params = string(res)
	}

	child := e.child.Error()
	if !strings.HasPrefix(child, "{") {
		child = fmt.Sprintf(`"%s"`, child)
	}

	return fmt.Sprintf(`{"operation":"%s", "description":"%s", "parameters":%s, "child":%s}`, e.operation, e.description, params, child)
}

func (e fault) Unwrap() error {
	return e.child
}
