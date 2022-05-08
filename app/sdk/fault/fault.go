package fault

import (
	"encoding/json"
	"errors"
	"fmt"
)

var (
	ErrNotFound            = errors.New("this resource could not be found or does not exist")
	ErrAlreadyExists       = errors.New("this resource could not be created because it already exists")
	ErrResourceHasChildren = errors.New("this resource could not be deleted because it has children")
	ErrMissingUID          = errors.New("this resource could not be found because the id provided is not valid")
)

type AdditionalData map[string]interface{}

type Error struct {
	Operation  string
	Err        error
	Parameters AdditionalData
	ParentErr  *Error
}

func (e Error) Error() string {

	var parameters string
	if e.Parameters != nil {

		res, err := json.Marshal(e.Parameters)
		if err == nil {
			parameters = string(res)
			parameters = parameters[1 : len(parameters)-1]
		}

		return fmt.Sprintf("%s with %s -> %s", e.Operation, parameters, e.Err.Error())
	}

	return fmt.Sprintf("%s -> %s", e.Operation, e.Err.Error())
}

func (e Error) Unwrap() error {
	return e.Err
}

func Wrap(err error, operation string, parameters ...AdditionalData) *Error {
	if err == nil {
		return nil
	}

	var prm AdditionalData
	if len(parameters) > 0 {
		prm = parameters[0]
	}

	domainError := Error{
		Parameters: prm,
		Operation:  operation,
		Err:        err,
	}

	if typedError, ok := err.(*Error); ok {
		typedError.ParentErr = &domainError
	}

	return &domainError
}
