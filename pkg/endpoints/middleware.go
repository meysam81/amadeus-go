package endpoints

import (
	"context"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func(begin time.Time) {
				_ = logger.Log(
					"layer", "endpoint",
					"input", request,
					"output", response,
					"error", err,
					"took", time.Since(begin),
				)
			}(time.Now())

			return next(ctx, request)
		}
	}
}
