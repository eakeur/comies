package handler

import (
	"comies/app/sdk/types"
	"context"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type (
	Routine func(ctx context.Context, r *http.Request) Response
	Route   struct {
		middlewares []string
		methods     []string
		path        string
		routine     interface{}
	}
)

func (r Route) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	if r.routine == nil {
		writer.WriteHeader(500)
		return
	}

	if callee, ok := r.routine.(func(ctx context.Context, r *http.Request) Response); ok {
		res := callee(request.Context(), request)
		res.Write(writer, request)
		return
	}

	if callee, ok := r.routine.(func(ctx context.Context, w http.ResponseWriter, r *http.Request) Response); ok {
		res := callee(request.Context(), writer, request)
		res.Write(writer, request)
		return
	}

	writer.WriteHeader(500)
}

func GetURLParam(r *http.Request, name string) string {
	return chi.URLParam(r, name)
}

func GetResourceIDFromURL(r *http.Request, name string) (types.ID, error) {
	v := chi.URLParam(r, name)
	return ConvertToID(v)
}

func ConvertToID(in string) (types.ID, error) {
	id, err := strconv.Atoi(in)
	if err != nil {
		return 0, err
	}

	return types.ID(id), nil
}
