package handler

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"reflect"
	"strings"
)

const (
	MethodTag                 = "method"
	MiddlewareTag             = "middleware"
	PatternTag                = "path"
	URLTag                    = "url"
	QueryParamsPlaceholderTag = "params"
	BodyPlaceholderTag        = "body"
)

type (
	Handler struct {
		middlewares map[string]Middleware
		Router      chi.Router
	}

	Middleware func(handler http.Handler) http.Handler

	Response interface {
		Write(w http.ResponseWriter, r *http.Request)
	}
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

		routeFieldValue.Set(reflect.ValueOf(Route{
			name:        routeFieldType.Name,
			methods:     strings.Split(tag.Get(MethodTag), ","),
			path:        tag.Get(PatternTag),
			bodyStruct:  tag.Get(BodyPlaceholderTag),
			query:       tag.Get(QueryParamsPlaceholderTag),
			urlParams:   tag.Get(URLTag),
			middlewares: strings.Split(tag.Get(MiddlewareTag), ","),
			routine:     handlerVal.MethodByName(method),
		}))

		ru := routeFieldValue.Interface().(Route)

		var middlewares []func(handler http.Handler) http.Handler
		for _, name := range ru.middlewares {
			if mid, ok := h.middlewares[name]; ok {
				middlewares = append(middlewares, mid)
			}
		}

		if len(middlewares) > 0 {
			router = router.With(middlewares...)
		}

		for _, m := range ru.methods {
			router.Method(m, ru.path, ru)
		}

	}

	return router

}
