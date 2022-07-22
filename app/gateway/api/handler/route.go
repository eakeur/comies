package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
	"reflect"
	"strings"
)

type Route struct {
	middlewares []string
	methods     []string
	path        string
	bodyStruct  string
	query       string
	name        string
	urlParams   string
	routine     reflect.Value
}

type RouteParams map[string]string

func (r Route) unmarshalJSON(reader io.ReadCloser, to reflect.Type) (reflect.Value, error) {
	instance := reflect.New(to)
	ptr := instance.Interface()

	err := json.NewDecoder(reader).Decode(&ptr)
	if err != nil {
		return reflect.Value{}, err
	}

	return instance.Elem(), nil
}

func (r Route) params(_ http.ResponseWriter, request *http.Request) ([]reflect.Value, error) {
	var (
		routeSignature     = r.routine.Type()
		numberOfParameters = routeSignature.NumIn()
		parameters         = make([]reflect.Value, numberOfParameters)
	)

	parameters[0] = reflect.ValueOf(request.Context())
	for i := 1; i < numberOfParameters; i++ {
		parameter := routeSignature.In(i)

		if parameter.AssignableTo(reflect.TypeOf(request)) {
			parameters[i] = reflect.ValueOf(request)
			continue
		}

		if q := request.URL.Query(); parameter.AssignableTo(reflect.TypeOf(q)) {
			parameters[i] = reflect.ValueOf(q)
			continue
		}

		if parameter.Name() == r.bodyStruct && parameter.Kind() == reflect.Struct {
			val, err := r.unmarshalJSON(request.Body, parameter)
			if err != nil {
				return nil, err
			}

			parameters[i] = val
			continue
		}

		params := RouteParams{}
		if parameter.AssignableTo(reflect.TypeOf(params)) {
			for _, v := range strings.Split(r.urlParams, ",") {
				params[v] = chi.URLParam(request, v)
			}
			parameters[i] = reflect.ValueOf(params)
			continue
		}
	}

	return parameters, nil
}

func (r Route) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	params, err := r.params(writer, request)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	res := r.routine.Call(params)[0].Interface().(ResponseWriter)
	res.Write(writer)

}
