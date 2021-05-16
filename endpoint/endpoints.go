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

	CreateHome 		endpoint.Endpoint
	GetHome 		endpoint.Endpoint
	DeleteHome 		endpoint.Endpoint
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

	var createHome endpoint.Endpoint
	{
		createHome = MakeCreateHomeEndpoint(s)
	}
	var getHome endpoint.Endpoint
	{
		getHome = MakeGetHomeEndpoint(s)
	}

	return Endpoints{
		CreateCrime: createCrime,
		GetCrimes: getCrimes,
		GetCrime: getCrime,
		UpdateCrime: updateCrime,
		DeleteCrime: deleteCrime,
		CreateHome: createHome,
		GetHome: getHome,
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

type GetCrimesResponse struct {
	Crimes []domain.Crime `json:"crimes"`
}

func MakeGetCrimesEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*GetCrimesRequest)
		res, err := s.GetCrimes(ctx)
		if err != nil {
			return nil, err
		}

		return GetCrimesResponse{*res}, nil
	}
}

type GetCrimeRequest struct {
	ID 		int64
}

type GetCrimeResponse struct {
	domain.Crime
}

func MakeGetCrimeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetCrimeRequest)
		res, err := s.GetCrime(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return GetCrimeResponse{*res}, nil
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

type CreateHomeRequest struct {
	ID        int64     `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	UserName  string  `json:"user_name"`
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
	Image     string  `json:"image"`
}

type CreateHomeResponse struct {
	domain.Home
}

func MakeCreateHomeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CreateHomeRequest)
		res, err := s.CreateHome(ctx, &domain.Home{
			ID:        req.ID,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			UserName:  req.UserName,
			Longitude: req.Longitude,
			Latitude:  req.Latitude,
			Image:     req.Image,
		})
		if err != nil {
			return nil, err
		}

		return CreateHomeResponse{*res}, nil
	}
}

type GetHomeRequest struct {
	ID 	int64
}

type GetHomeResponse struct {
	domain.Home
}

func MakeGetHomeEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetHomeRequest)
		res, err := s.GetHome(ctx, req.ID)
		if err != nil {
			return nil, err
		}

		return GetHomeResponse{*res}, nil
	}
}