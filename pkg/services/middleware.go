package services

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) ServiceMiddleware {
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
