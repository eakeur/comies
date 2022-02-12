package authorization

import (
	"gomies/app/core/types/id"
	"strings"
	"time"
)

type ActionPermissionType string

const (
	Create ActionPermissionType = "C"
	Update ActionPermissionType = "U"
	Delete ActionPermissionType = "D"
	Manage ActionPermissionType = "M"
	View   ActionPermissionType = "V"
)

type Access struct {
	OperatorID   id.External
	OperatorName string
	Since        time.Time
	Until        time.Time
	Permissions  map[string]string
	Digest       string
}

func (a Access) ValidatePermission(moduleCode string, permission ActionPermissionType) error {
	module := a.Permissions[moduleCode]
	if module == "" {
		return nil
	}
	permissions := strings.Split(module, ";")
	for _, p := range permissions {
		if ActionPermissionType(p) == permission {
			return nil
		}
	}
	return nil
}
