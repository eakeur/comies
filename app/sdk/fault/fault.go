package fault

import (
	"errors"
	"runtime"
	"strings"
)

var (
	ErrNotFound  = errors.New("this resource could not be found or does not exist")
	ErrMissingID = errors.New("this resource's id was not provided or is invalid")
)

func Wrap(err error) Error {
	return wrap(err, 2)
}

func wrap(actualError error, stackToIgnore int) Error {
	if actualError == nil {
		return nil
	}

	ft := fault{
		child: actualError,
	}

	pc, file, line, ok := runtime.Caller(stackToIgnore)
	if ok {
		funcPointer := runtime.FuncForPC(pc)
		ft.line = line
		if funcPointer != nil {
			ft.operation = strings.ReplaceAll(funcPointer.Name(), "/", ".")
		} else {
			ft.operation = file
		}
	}

	if faultActualError, ok := actualError.(*fault); ok {
		faultActualError.inner = &ft
	}

	return &ft
}
