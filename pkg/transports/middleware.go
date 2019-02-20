package transports

import (
	pbFunc "amadeus-go/api/amadeus/func"
	pbType "amadeus-go/api/amadeus/type"

	"context"
	"time"

	"github.com/go-kit/kit/log"
)

func loggingMiddleware(logger log.Logger) TransportMiddleware {
	return func(next pbFunc.AmadeusServiceServer) pbFunc.AmadeusServiceServer {
		return logmw{logger, next}
	}
}

type logmw struct {
	logger log.Logger
	sv     pbFunc.AmadeusServiceServer
}

func (mw logmw) FlightLowFareSearch(ctx context.Context, req *pbFunc.FlightLowFareSearchRequest) (resp *pbType.FlightLowFareSearchResult, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"layer", "transport",
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
