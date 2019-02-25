package services

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) serviceMiddleware {
	return func(next AmadeusService) AmadeusService {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	sv     AmadeusService
}

func (mw logmw) FlightLowFareSearch(ctx context.Context, req *FlightLowFareSearchRequest) (resp *FlightLowFareSearchResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "FlightLowFareSearch",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.FlightLowFareSearch(ctx, req)
	return
}

func (mw logmw) FlightInspirationSearch(ctx context.Context, req *FlightInspirationSearchRequest) (resp *FlightInspirationSearchResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "FlightInspirationSearch",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.FlightInspirationSearch(ctx, req)
	return
}

func (mw logmw) FlightMostTraveledDestinations(ctx context.Context, req *FlightMostTraveledDestinationsRequest) (resp *FlightMostTraveledDestinationsResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "FlightMostTraveledDestinations",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.FlightMostTraveledDestinations(ctx, req)
	return
}

func (mw logmw) FlightMostBookedDestinations(ctx context.Context, req *FlightMostBookedDestinationsRequest) (resp *FlightMostBookedDestinationsResponse, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "FlightMostBookedDestinations",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.FlightMostBookedDestinations(ctx, req)
	return
}
