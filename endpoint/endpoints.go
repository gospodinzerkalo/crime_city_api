package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/gospodinzerkalo/crime_city_api/domain"
	"github.com/gospodinzerkalo/crime_city_api/service"
	"log"
)

type Endpoints struct {
	CreateCrime 	endpoint.Endpoint
	GetCrimes	 	endpoint.Endpoint
	UpdateCrime 	endpoint.Endpoint
	DeleteCrime 	endpoint.Endpoint
}

func NewEndpointsFactory(s service.Service, log log.Logger) Endpoints {
	var createCrime endpoint.Endpoint
	{
		createCrime = MakeCreateCrimeEndpoint(s)
	}

	return Endpoints{
		CreateCrime: createCrime,
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