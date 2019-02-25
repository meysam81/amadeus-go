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
	FlightMostBookedDestinationsEndpoint   endpoint.Endpoint
	FlightBusiestTravelingPeriodEndpoint   endpoint.Endpoint
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

func (s AmadeusEndpointSet) FlightMostBookedDestinations(ctx context.Context, request *sv.FlightMostBookedDestinationsRequest) (*sv.FlightMostBookedDestinationsResponse, error) {
	resp, err := s.FlightMostBookedDestinationsEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightMostBookedDestinationsResponse)
	return response, nil
}

func (s AmadeusEndpointSet) FlightBusiestTravelingPeriod(ctx context.Context, request *sv.FlightBusiestTravelingPeriodRequest) (*sv.FlightBusiestTravelingPeriodResponse, error) {
	resp, err := s.FlightBusiestTravelingPeriodEndpoint(ctx, request)
	if err != nil {
		return nil, err
	}

	response := resp.(*sv.FlightBusiestTravelingPeriodResponse)
	return response, nil
}

func NewEndpointSet(srv sv.AmadeusService, logger log.Logger) *AmadeusEndpointSet {
	var flightLowFareSearchEndpoint endpoint.Endpoint
	flightLowFareSearchEndpoint = makeFlightLowFareSearchEndpoint(srv)
	flightLowFareSearchEndpoint = loggingMiddleware(logger, "FlightLowFareSearch")(flightLowFareSearchEndpoint)

	var flightInspirationSearchEndpoint endpoint.Endpoint
	flightInspirationSearchEndpoint = makeFlightInspirationSearchEndpoint(srv)
	flightInspirationSearchEndpoint = loggingMiddleware(logger, "FlightInspirationSearch")(flightInspirationSearchEndpoint)

	var flightMostTraveledDestinationsEndpoint endpoint.Endpoint
	flightMostTraveledDestinationsEndpoint = makeFlightMostTraveledDestinationsEndpoint(srv)
	flightMostTraveledDestinationsEndpoint = loggingMiddleware(logger, "FlightMostTraveledDestinations")(flightMostTraveledDestinationsEndpoint)

	var flightMostBookedDestinationsEndpoint endpoint.Endpoint
	flightMostBookedDestinationsEndpoint = makeFlightMostBookedDestinationsEndpoint(srv)
	flightMostBookedDestinationsEndpoint = loggingMiddleware(logger, "FlightMostBookedDestinations")(flightMostBookedDestinationsEndpoint)

	return &AmadeusEndpointSet{
		FlightLowFareSearchEndpoint:            flightLowFareSearchEndpoint,
		FlightInspirationSearchEndpoint:        flightInspirationSearchEndpoint,
		FlightMostTraveledDestinationsEndpoint: flightMostTraveledDestinationsEndpoint,
		FlightMostBookedDestinationsEndpoint:   flightMostBookedDestinationsEndpoint,
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

func makeFlightMostBookedDestinationsEndpoint(srv sv.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*sv.FlightMostBookedDestinationsRequest)
		if !ok {
			return nil, errors.New("service did not fetch type <FlightMostBookedDestinationsRequest>")
		}

		resp, err := srv.FlightMostBookedDestinations(ctx, req)
		return resp, err
	}
}
