package types

import (
	"errors"
	"strings"
)

type Permissions string
var ErrNotAllowed = errors.New("the actual session cannot operate the targeted function")
func (p Permissions) CheckAllowance(operation string) error {
	if p == "*" {
		return nil
	}

	if !strings.Contains(string(p), operation) {
		return ErrNotAllowed
	}

	return nil
}
