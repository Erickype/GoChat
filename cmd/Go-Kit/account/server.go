package account

import (
	"context"
	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func NewHttpServer(_ context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/user").Handler(httpTransport.NewServer(
		endpoints.CreateUser,
		decodeUserRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httpTransport.NewServer(
		endpoints.GetUser,
		decodeEmailRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
