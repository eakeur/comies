package send

import (
	"comies/app/core/types"
	"errors"
	"net/http"
)

func FromError(err error, options ...Option) Response {
	response, ok := failures[err]
	if !ok {
		response.Code = 500
		response.Data = InternalServerError
		for mappedError, mappedResponse := range failures {
			if errors.Is(err, mappedError) {
				response = mappedResponse
				break
			}
		}
	}

	return build(response, append(options, WithError(err))...)
}

func Data(code int, data interface{}, options ...Option) Response {
	return build(Response{Data: data, Code: code}, options...)
}

func Status(code int, options ...Option) Response {
	return Data(code, nil, options...)
}

func CreatedWithID(id types.ID) Response {
	return Status(http.StatusCreated, WithHeaders(map[string]string{"Location": id.String()}))
}

func build(r Response, options ...Option) Response {
	if len(options) == 0 {
		return r
	}

	return build(options[0](r), options[1:]...)
}
