package endpoints

import (
	sv "amadeus-go/pkg/services"

	"context"
	"errors"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

type AmadeusEndpointSet struct {
	FlightLowFareSearchEndpoint            endpoint.Endpoint
	FlightInspirationSearchEndpoint        endpoint.Endpoint
	FlightMostTraveledDestinationsEndpoint endpoint.Endpoint
}

func (s AmadeusEndpointSet) FlightLowFareSearch(ctx context.Context, request *sv.FlightLowFareSearchRequest) (*sv.FlightLowFareSearchResponse, error) {
	resp, err := s.FlightLowFareSearchEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightLowFareSearchResponse)
	return response, nil
}

func (s AmadeusEndpointSet) FlightInspirationSearch(ctx context.Context, request *sv.FlightInspirationSearchRequest) (*sv.FlightInspirationSearchResponse, error) {
	resp, err := s.FlightInspirationSearchEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightInspirationSearchResponse)
	return response, nil
}

func (s AmadeusEndpointSet) FlightMostTraveledDestinations(ctx context.Context, request *sv.FlightMostTraveledDestinationsRequest) (*sv.FlightMostTraveledDestinationsResponse, error) {
	resp, err := s.FlightMostTraveledDestinationsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightMostTraveledDestinationsResponse)
	return response, nil
}

func NewEndpointSet(srv sv.AmadeusService, logger log.Logger) *AmadeusEndpointSet {
	var flightLowFareSearchEndpoint endpoint.Endpoint
	flightLowFareSearchEndpoint = makeFlightLowFareSearchEndpoint(srv)
	flightLowFareSearchEndpoint = loggingMiddleware(logger, "FlightLowFareSearch")(flightLowFareSearchEndpoint)

	var flightInspirationSearchEndpoint endpoint.Endpoint
	flightInspirationSearchEndpoint = makeFlightInspirationSearchEndpoint(srv)
	flightInspirationSearchEndpoint = loggingMiddleware(logger, "FlightInspirationSearch")(flightInspirationSearchEndpoint)

	var flightCheapestDateSearchEndpoint endpoint.Endpoint
	flightCheapestDateSearchEndpoint = makeFlightMostTraveledDestinationsEndpoint(srv)
	flightCheapestDateSearchEndpoint = loggingMiddleware(logger, "FlightMostTraveledDestinations")(flightCheapestDateSearchEndpoint)

	return &AmadeusEndpointSet{
		FlightLowFareSearchEndpoint:            flightLowFareSearchEndpoint,
		FlightInspirationSearchEndpoint:        flightInspirationSearchEndpoint,
		FlightMostTraveledDestinationsEndpoint: flightCheapestDateSearchEndpoint,
	}
}

// ====================================================================================================
func makeFlightLowFareSearchEndpoint(srv sv.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*sv.FlightLowFareSearchRequest)
		if !ok {
			return nil, errors.New("service did not fetch type <FlightLowFareSearchRequest>")
		}

		resp, err := srv.FlightLowFareSearch(ctx, req)
		return resp, err
	}
}

func makeFlightInspirationSearchEndpoint(srv sv.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*sv.FlightInspirationSearchRequest)
		if !ok {
			return nil, errors.New("service did not fetch type <FlightInspirationSearchRequest>")
		}

		resp, err := srv.FlightInspirationSearch(ctx, req)
		return resp, err
	}
}

func makeFlightMostTraveledDestinationsEndpoint(srv sv.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*sv.FlightMostTraveledDestinationsRequest)
		if !ok {
			return nil, errors.New("service did not fetch type <FlightMostTraveledDestinationsRequest>")
		}

		resp, err := srv.FlightMostTraveledDestinations(ctx, req)
		return resp, err
	}
}
