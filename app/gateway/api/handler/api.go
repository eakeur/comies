package handler

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"reflect"
	"strings"
)

const (
	MethodTag     = "method"
	MiddlewareTag = "middleware"
	PatternTag    = "path"
)

type (
	Handler struct {
		middlewares map[string]Middleware
		Router      chi.Router
	}

	Middleware func(handler http.Handler) http.Handler
)

func NewHandler(middlewares map[string]Middleware) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}

func (h *Handler) RegisterService(router chi.Router, service interface{}) chi.Router {

	handlerVal := reflect.ValueOf(service).Elem()
	handlerType := handlerVal.Type()

	for i := 0; i < handlerType.NumField(); i++ {
		routeFieldType := handlerType.Field(i)
		routeFieldValue := handlerVal.Field(i)

		if routeFieldValue.Type().Name() != "Route" {
			continue
		}

		tag := routeFieldType.Tag
		method := strings.TrimSuffix(routeFieldType.Name, "Route")

		routine := handlerVal.MethodByName(method)
		if !routine.IsValid() {
			log.Fatalf("Could not find a routine to run for route %v. The routine and route name must be strictly the same", method)
		}

		routeFieldValue.Set(reflect.ValueOf(Route{
			methods:     strings.Split(tag.Get(MethodTag), ","),
			path:        tag.Get(PatternTag),
			middlewares: strings.Split(tag.Get(MiddlewareTag), ","),
			routine:     routine.Interface(),
		}))

		ru := routeFieldValue.Interface().(Route)

		var middlewares []func(handler http.Handler) http.Handler
		for _, name := range ru.middlewares {
			if mid, ok := h.middlewares[name]; ok {
				middlewares = append(middlewares, mid)
			}
		}

		r := router
		if len(middlewares) > 0 {
			r = r.With(middlewares...)
		}

		for _, m := range ru.methods {
			r.Method(m, ru.path, ru)
		}

	}

	return router

}
