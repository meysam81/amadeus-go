package endpoints

import (
	sv "Amadeus/amadeus-go/pkg/services"
	"github.com/go-kit/kit/log"
	"os"

	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type AmadeusEndpointSet struct {
	FlightLowFareSearchEndpoint endpoint.Endpoint
}

func (s AmadeusEndpointSet) FlightLowFareSearch(ctx context.Context, routeData *sv.FlightLowFareSearchRequest) (*sv.FlightLowFareSearchResponse, error) {
	resp, err := s.FlightLowFareSearchEndpoint(ctx, routeData)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightLowFareSearchResponse)
	return response, nil
}

func MakeFlightLowFareSearchEndpoint(srv sv.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*sv.FlightLowFareSearchRequest)
		if !ok {
			return nil, errors.New("bad request")
		}
		resp, err := srv.FlightLowFareSearch(ctx, req)
		return resp, err
	}
}

func NewEndpointSet(srv sv.AmadeusService) AmadeusEndpointSet {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller)

	var flightLowFareSearchEndpoint endpoint.Endpoint
	flightLowFareSearchEndpoint = MakeFlightLowFareSearchEndpoint(srv)

	flightLowFareSearchEndpoint = loggingMiddleware(logger)(flightLowFareSearchEndpoint)

	return AmadeusEndpointSet{
		FlightLowFareSearchEndpoint: flightLowFareSearchEndpoint,
	}
}
