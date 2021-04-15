package phonebook

import (
	"context"
	"encoding/json"
	"net/http"

	endpoint "github.com/sapawarga/api-orchestration/endpoint/phonebook"
	"github.com/sapawarga/api-orchestration/helper"
	"github.com/sapawarga/api-orchestration/usecase"

	kitlog "github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type err interface {
	error() error
}

func MakeHandler(ctx context.Context, fs usecase.PhonebookI, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	processGetListPhonebook := kithttp.NewServer(
		endpoint.MakeGetList(ctx, fs),
		decodeGetListRequest,
		encodeResponse,
		opts...,
	)

	processGetDetailPhonebook := kithttp.NewServer(
		endpoint.MakeGetDetail(ctx, fs),
		decodeGetDetailRequest,
		encodeResponse,
		opts...,
	)

	processCreatePhonebook := kithttp.NewServer(
		endpoint.MakeCreatePhonebook(ctx, fs),
		decodeCreateRequest,
		encodeResponse,
		opts...,
	)

	processUpdatePhonebook := kithttp.NewServer(
		endpoint.MakeUpdatePhonebook(ctx, fs),
		decodeUpdateRequest,
		encodeResponse,
		opts...,
	)

	processDeletePhonebook := kithttp.NewServer(
		endpoint.MakeDeletePhonebook(ctx, fs),
		decodeGetDetailRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/phonebooks", processGetListPhonebook).Methods(helper.GET)
	r.Handle("/phonebooks/{id}", processGetDetailPhonebook).Methods(helper.GET)
	r.Handle("/phonebooks", processCreatePhonebook).Methods(helper.POST)
	r.Handle("/phonebooks/{id}", processUpdatePhonebook).Methods(helper.PUT)
	r.Handle("/phonebooks/{id}", processDeletePhonebook).Methods(helper.DELETE)

	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(err); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusBadRequest)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})

}
