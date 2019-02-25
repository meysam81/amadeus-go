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

func (mw logmw) FlightLowFareSearch(ctx context.Context, req *FlightLowFareSearchRequest) (resp *Response, err error) {
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

func (mw logmw) FlightInspirationSearch(ctx context.Context, req *FlightInspirationSearchRequest) (resp *Response, err error) {
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

func (mw logmw) FlightMostTraveledDestinations(ctx context.Context, req *FlightMostTraveledDestinationsRequest) (resp *Response, err error) {
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

func (mw logmw) FlightMostBookedDestinations(ctx context.Context, req *FlightMostBookedDestinationsRequest) (resp *Response, err error) {
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

func (mw logmw) FlightBusiestTravelingPeriod(ctx context.Context, req *FlightBusiestTravelingPeriodRequest) (resp *Response, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "FlightBusiestTravelingPeriod",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.FlightBusiestTravelingPeriod(ctx, req)
	return
}

func (mw logmw) AirportNearestRelevant(ctx context.Context, req *AirportNearestRelevantRequest) (resp *Response, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "AirportNearestRelevant",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.AirportNearestRelevant(ctx, req)
	return
}

func (mw logmw) AirportAndCitySearch(ctx context.Context, req *AirportAndCitySearchRequest) (resp *Response, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "service",
			"method", "AirportAndCitySearch",
			"input", req,
			"output", resp,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	resp, err = mw.sv.AirportAndCitySearch(ctx, req)
	return
}

