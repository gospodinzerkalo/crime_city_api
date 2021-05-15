package transport

import (
	"context"
	"encoding/json"
	kitEndpoint "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"net/http"
)

func MakeHTTPHandler(endpoints endpoint.Endpoints, log log.Logger)  {
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(log)),
		httptransport.ServerErrorEncoder(errorEncoder),
	}

	createCrime := httptransport.NewServer(
		endpoints.CreateCrime,
		decodeHTTPCreateCrimeRequest,
		encodeHTTPGenericResponse,
		options...,
		)

	r.Handle("/crime", createCrime).Methods("POST")
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(err)
}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if f, ok := response.(kitEndpoint.Failer); ok && f.Failed() != nil {
		errorEncoder(ctx, f.Failed(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func decodeHTTPCreateCrimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &endpoint.CreateCrimeRequest{}

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return nil, err
	}

	return req, nil
}