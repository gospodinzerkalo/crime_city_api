package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/gospodinzerkalo/crime_city_api/domain"
	"github.com/gospodinzerkalo/crime_city_api/service"
)

type Endpoints struct {
	CreateCrime 	endpoint.Endpoint
	GetCrimes	 	endpoint.Endpoint
	GetCrime	 	endpoint.Endpoint
	UpdateCrime 	endpoint.Endpoint
	DeleteCrime 	endpoint.Endpoint
}

func NewEndpointsFactory(s service.Service, log log.Logger) Endpoints {
	var createCrime endpoint.Endpoint
	{
		createCrime = MakeCreateCrimeEndpoint(s)
	}
	var getCrimes endpoint.Endpoint
	{
		getCrimes = MakeGetCrimesEndpoint(s)
	}
	var getCrime endpoint.Endpoint
	{
		getCrime = MakeGetCrimeEndpoint(s)
	}
	var updateCrime endpoint.Endpoint
	{
		updateCrime = MakeUpdateCrimeEndpoint(s)
	}
	var deleteCrime endpoint.Endpoint
	{
		deleteCrime = MakeDeleteCrimeEndpoint(s)
	}

	return Endpoints{
		CreateCrime: createCrime,
		GetCrimes: getCrimes,
		GetCrime: getCrime,
		UpdateCrime: updateCrime,
		DeleteCrime: deleteCrime,
	}
}

type CreateCrimeRequest struct {
	domain.Crime
}

func MakeCreateCrimeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateCrimeRequest)

		return s.CreateCrime(ctx, &domain.Crime{
			LocationName: req.LocationName,
			Longitude:    req.Longitude,
			Latitude:     req.Latitude,
			Description:  req.Description,
			Image:        req.Image,
		})
	}
}

type GetCrimesRequest struct {

}

type GetCrimeResponse struct {
	Crimes []domain.Crime
}

func MakeGetCrimesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*GetCrimesRequest)
		return s.GetCrimes(ctx)
	}
}

type GetCrimeRequest struct {
	ID 		int64
}

func MakeGetCrimeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetCrimeRequest)
		return s.GetCrime(ctx, req.ID)
	}
}

type UpdateCrimeRequest struct {
	domain.Crime
}

func MakeUpdateCrimeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateCrimeRequest)
		return s.UpdateCrime(ctx, &domain.Crime{
			ID:           req.ID,
			LocationName: req.LocationName,
			Longitude:    req.Longitude,
			Latitude:     req.Latitude,
			Description:  req.Description,
			Image:        req.Image,
		})
	}
}

type DeleteCrimeRequest struct {
	ID 		int64
}

func MakeDeleteCrimeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*DeleteCrimeRequest)
		return nil, s.DeleteCrime(ctx, req.ID)
	}
}