package transports

import (
	"amadeus-go/pkg/endpoints"
	srv "amadeus-go/pkg/services"
	pbFunc "api/amadeus/func"
	pbType "api/amadeus/type"

	"context"
	"github.com/go-kit/kit/endpoint"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) srv.AmadeusService {
	var flightLowFareSearchEndpoint endpoint.Endpoint
	{
		flightLowFareSearchEndpoint = grpcTransport.NewClient(
			conn,
			"amadeus.func.AmadeusService",
			"FlightLowFareSearch",
			encodeRequest,
			decodeResponse,
			pbType.FlightLowFareSearchResult{},
		).Endpoint()
	}

	return endpoints.AmadeusEndpointSet{FlightLowFareSearchEndpoint: flightLowFareSearchEndpoint}
}

func encodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*srv.FlightLowFareSearchRequest)
	return &pbFunc.FlightLowFareSearchRequest{
		Origin:        req.Origin,
		DepartureDate: req.DepartureDate,
		Destination:   req.Destination,
		ReturnDate:    req.ReturnDate,
	}, nil
}

func decodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pbType.FlightLowFareSearchResult)
	return resp, nil
}
