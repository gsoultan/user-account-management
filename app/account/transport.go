package account

import (
	"context"
	"encoding/json"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

type errorer interface {
	error() error
}

func MakeHttpHandler(c CommandService, t TransactionService, logger kitlog.Logger) http.Handler {
	r := mux.NewRouter()

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(kitlog.Logger(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	signUpHandler := kithttp.NewServer(makeSignUpEndpoint(c), decodeSignUpRequest, encodeResponse, opts...)

	r.Handle("/account/sign-up/", signUpHandler).Methods("POST")
	return r
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeSignUpRequest(ctx context.Context, r *http.Request) (request interface{}, err error) {
	signUpRequest := SignUpRequest{}
	err = json.NewDecoder(r.Body).Decode(&signUpRequest)
	return signUpRequest, err
}

func encodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
