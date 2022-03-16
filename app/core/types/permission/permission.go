package permission

import (
	"strings"
)

type Permissions string

func (p Permissions) CheckAllowance(operation string) error {
	if p == "*" {
		return nil
	}

	if !strings.Contains(string(p), operation) {
		return ErrNotAllowed
	}

	return nil
}
