package transport

import (
	"context"
	"database/sql"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/gospodinzerkalo/crime_city_api/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCServer struct {
	getCrimes 		gt.Handler
	getCrime 		gt.Handler

	createHome		gt.Handler
	getHome			gt.Handler
	deleteHome 		gt.Handler
}

// NewGRPCServer initializes a new gRPC server
func NewGRPCServer(endpoints endpoint.Endpoints, logger log.Logger) pb.CrimeServiceServer {
	return &gRPCServer{
		getCrimes: gt.NewServer(
			endpoints.GetCrimes,
			decodeGetCrimesRequest,
			encodeGetCrimesResponse,
		),
		getCrime: gt.NewServer(
			endpoints.GetCrime,
			decodeGetCrimeRequest,
			encodeGetCrimeResponse,
		),
		createHome: gt.NewServer(
			endpoints.CreateHome,
			decodeCreateHomeRequest,
			encodeCreateHomeResponse,
		),
		getHome: gt.NewServer(
			endpoints.GetHome,
			decodeGetHomeRequest,
			encodeGetHomeResponse,
		),
		deleteHome: gt.NewServer(
			endpoints.DeleteHome,
			decodeDeleteHomeRequest,
			encodeDeleteHomeResponse,
			),
	}
}

func grpcErrorEncoder(err error) error {
	switch err {
	case sql.ErrNoRows:
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}

func (g *gRPCServer) GetCrimes(ctx context.Context, request *pb.GetCrimesRequest) (*pb.GetCrimesResponse, error) {
	_, resp, err := g.getCrimes.ServeGRPC(ctx, request)
	if err != nil {
		return nil, grpcErrorEncoder(err)
	}
	return resp.(*pb.GetCrimesResponse), nil
}

func decodeGetCrimesRequest(_ context.Context, request interface{}) (interface{}, error) {
	_ = request.(*pb.GetCrimesRequest)
	return &endpoint.GetCrimesRequest{}, nil
}

func encodeGetCrimesResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetCrimesResponse)
	var crimes []*pb.Crime
	for _, crime := range resp.Crimes {
		crimes = append(crimes, &pb.Crime{
			Id:           crime.ID,
			LocationName: crime.LocationName,
			Longitude:    crime.Longitude,
			Latitude:     crime.Latitude,
			Description:  crime.Description,
			Image:        crime.Image,
			Date:         crime.Date,
		})
	}
	return &pb.GetCrimesResponse{Crimes: crimes}, nil
}

func (g *gRPCServer) GetCrime(ctx context.Context, request *pb.GetCrimeRequest) (*pb.GetCrimeResponse, error) {
	_, resp, err := g.getCrime.ServeGRPC(ctx, request)
	if err != nil {
		return nil, grpcErrorEncoder(err)
	}
	return resp.(*pb.GetCrimeResponse), nil
}

func decodeGetCrimeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetCrimeRequest)
	return &endpoint.GetCrimeRequest{ID: req.Id}, nil
}

func encodeGetCrimeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetCrimeResponse)

	return &pb.GetCrimeResponse{Crime: &pb.Crime{
		Id:           resp.ID,
		LocationName: resp.LocationName,
		Longitude:    resp.Longitude,
		Latitude:     resp.Latitude,
		Description:  resp.Description,
		Image:        resp.Image,
		Date:         resp.Date,
	}}, nil
}

func (g *gRPCServer) CreateHome(ctx context.Context, request *pb.CreateHomeRequest) (*pb.CreateHomeResponse, error) {
	_, resp, err := g.createHome.ServeGRPC(ctx, request)
	if err != nil {
		return nil, grpcErrorEncoder(err)
	}
	return resp.(*pb.CreateHomeResponse), nil
}

func decodeCreateHomeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.CreateHomeRequest)
	return &endpoint.CreateHomeRequest{
		ID:        req.Home.Id,
		FirstName: req.Home.FirstName,
		LastName:  req.Home.LastName,
		UserName:  req.Home.UserName,
		Longitude: req.Home.Longitude,
		Latitude:  req.Home.Latitude,
		Image:     req.Home.Image,
	}, nil
}

func encodeCreateHomeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.CreateHomeResponse)
	return &pb.CreateHomeResponse{Home: &pb.Home{
		Id:        resp.ID,
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		UserName:  resp.UserName,
		Longitude: resp.Longitude,
		Latitude:  resp.Latitude,
		Image:     resp.Image,
	}}, nil
}

func (g *gRPCServer) GetHome(ctx context.Context, request *pb.GetHomeRequest) (*pb.GetHomeResponse, error) {
	_, resp, err := g.getHome.ServeGRPC(ctx, request)
	if err != nil {
		return nil, grpcErrorEncoder(err)
	}
	return resp.(*pb.GetHomeResponse), nil
}

func decodeGetHomeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.GetHomeRequest)
	return &endpoint.GetHomeRequest{ID: req.Id}, nil
}

func encodeGetHomeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoint.GetHomeResponse)
	return &pb.GetHomeResponse{Home: &pb.Home{
		Id:        resp.ID,
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		UserName:  resp.UserName,
		Longitude: resp.Longitude,
		Latitude:  resp.Latitude,
		Image:     resp.Image,
	}}, nil
}

func (g *gRPCServer) DeleteHome(ctx context.Context, request *pb.DeleteHomeRequest) (*pb.DeleteHomeResponse, error) {
	_, resp, err := g.deleteHome.ServeGRPC(ctx, request)
	if err != nil {
		return nil, grpcErrorEncoder(err)
	}
	return resp.(*pb.DeleteHomeResponse), nil
}

func decodeDeleteHomeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*pb.DeleteHomeRequest)
	return &endpoint.DeleteHomeRequest{ID: req.Id}, nil
}

func encodeDeleteHomeResponse(_ context.Context, response interface{}) (interface{}, error) {
	_ = response.(endpoint.DeleteHomeResponse)
	return &pb.DeleteHomeResponse{}, nil
}
