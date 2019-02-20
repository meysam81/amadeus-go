package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			defer func(begin time.Time) {
				_ = logger.Log(
					"layer", "endpoint",
					"method", "FlightLowFareSearch",
					"input", req,
					"output", resp,
					"error", err,
					"took", time.Since(begin),
				)
			}(time.Now())
			return next(ctx, req)
		}
	}
}
