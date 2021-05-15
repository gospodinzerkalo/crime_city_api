package transport

import (
	"context"
	"github.com/go-kit/kit/log"
	gt "github.com/go-kit/kit/transport/grpc"
	"github.com/gospodinzerkalo/crime_city_api/endpoint"
	"github.com/gospodinzerkalo/crime_city_api/pb"
)

type gRPCServer struct {
	getCrimes 		gt.Handler
	getCrime 		gt.Handler
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

	}
}

func (g *gRPCServer) GetCrimes(ctx context.Context, request *pb.GetCrimesRequest) (*pb.GetCrimesResponse, error) {
	_, resp, err := g.getCrimes.ServeGRPC(ctx, request)
	if err != nil {
		return nil, err
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
		return nil, err
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