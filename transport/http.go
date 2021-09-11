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
	"github.com/gospodinzerkalo/crime_city_api/service"
	"net/http"
	"strconv"
)

func MakeHTTPHandler(endpoints endpoint.Endpoints, log log.Logger) http.Handler  {
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

	getCrimes := httptransport.NewServer(
		endpoints.GetCrimes,
		decodeHTTPGetCrimesRequest,
		encodeHTTPGenericResponse,
		options...,
		)

	getCrime := httptransport.NewServer(
		endpoints.GetCrime,
		decodeHTTPGetCrimeRequest,
		encodeHTTPGenericResponse,
		options...,
		)

	updateCrime := httptransport.NewServer(
		endpoints.UpdateCrime,
		decodeHTTPUpdateCrimeRequest,
		encodeHTTPGenericResponse,
		options...,
		)

	deleteCrime := httptransport.NewServer(
		endpoints.DeleteCrime,
		decodeHTTPDeleteCrimeRequest,
		encodeHTTPGenericResponse,
		options...,
		)

	r.Handle("/crime", createCrime).Methods("POST")
	r.Handle("/crimes", getCrimes).Methods("GET")
	r.Handle("/crime/{id}", getCrime).Methods("GET")
	r.Handle("/crime/{id}", updateCrime).Methods("PUT")
	r.Handle("/crime/{id}", deleteCrime).Methods("DELETE")

	return r
}

func errorEncoder(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case service.ErrIdIncorrect, service.ErrInvalidData, service.ErrInvalidPathParam:
		w.WriteHeader(http.StatusBadRequest)
	case service.ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	}
	json.NewEncoder(w).Encode(err.Error())
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
		return nil, service.ErrInvalidData
	}

	return req, nil
}

func decodeHTTPGetCrimesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := &endpoint.GetCrimesRequest{}
	return req, nil
}

func decodeHTTPGetCrimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := extractIdMux(r)
	if err != nil {
		return nil, err
	}
	req := &endpoint.GetCrimeRequest{ID: id}
	return req, nil
}

func decodeHTTPUpdateCrimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id, err := extractIdMux(r)
	if err != nil {
		return nil, err
	}
	req := &endpoint.UpdateCrimeRequest{}
	req.ID = id

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func decodeHTTPDeleteCrimeRequest(_ context.Context, r *http.Request) (interface{}, error) {
	id , err := extractIdMux(r)
	if err != nil {
		return nil, err
	}
	req := &endpoint.DeleteCrimeRequest{ID: id}
	return req, nil
}

func extractIdMux(r *http.Request) (int64, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return 0, service.ErrIdIncorrect
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}

	return int64(intId), nil
}
