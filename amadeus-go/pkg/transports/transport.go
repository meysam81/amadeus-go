package transports

import (
	"amadeus-go/pkg/endpoints"
	sv "amadeus-go/pkg/services"
	pbFunc "api/amadeus/func"
	pbType "api/amadeus/type"
	//pbComn "api/amadeus/comn"

	"github.com/go-kit/kit/log"
	"os"

	"context"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	FlightLowFareSearchHandler grpcTransport.Handler
}

func (s *grpcServer) FlightLowFareSearch(ctx context.Context, req *pbFunc.FlightLowFareSearchRequest) (*pbType.FlightLowFareSearchResult, error) {
	_, resp, err := s.FlightLowFareSearchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.FlightLowFareSearchResult)
	return response, nil
}

func NewGRPCServer(endpoints endpoints.AmadeusEndpointSet) (s pbFunc.AmadeusServiceServer) {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller)

	g := &grpcServer{
		FlightLowFareSearchHandler: grpcTransport.NewServer(
			endpoints.FlightLowFareSearchEndpoint,
			decodeFlightLowFareSearchRequest,
			encodeFlightLowFareSearchResponse,
		),
	}

	s = loggingMiddleware(logger)(g)
	return
}

func decodeFlightLowFareSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pbFunc.FlightLowFareSearchRequest)
	return &sv.FlightLowFareSearchRequest{
		Origin:        req.Origin,
		DepartureDate: req.DepartureDate,
		Destination:   req.Destination,
		ReturnDate:    req.ReturnDate,
	}, nil
}

func encodeFlightLowFareSearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(sv.Data)

	var off []*pbType.OfferItem
	for _, offers := range resp.OfferItems {

		var services []*pbType.Service
		for _, service := range offers.Services {

			var segments []*pbType.Segment
			for _, segment := range service.Segments {

				var flightSegments []*pbType.FlightSegment
				for _, flightSegment := range segment.FlightSegment {
					flightSegments = append(flightSegments, &pbType.FlightSegment{
						Duration: flightSegment.Duration,
						Number:   flightSegment.Number,
						Aircraft: &pbType.Aircraft{
							Code: flightSegment.Aircraft.Code,
						},
						Arrival: &pbType.DepartureArrival{
							At:       flightSegment.Arrival.At,
							IataCode: flightSegment.Arrival.IataCode,
							Terminal: flightSegment.Arrival.Terminal,
						},
						Departure: &pbType.DepartureArrival{
							At:       flightSegment.Departure.At,
							IataCode: flightSegment.Departure.IataCode,
							Terminal: flightSegment.Departure.Terminal,
						},
						CarrierCode: flightSegment.CarrierCode,
						Operating: &pbType.Operating{
							CarrierCode: flightSegment.Operating.CarrierCode,
							Number:      flightSegment.Operating.Number,
						},
					})
				}

				segments = append(segments, &pbType.Segment{
					FlightSegments: flightSegments,
					PricingDetailPerAdult: &pbType.PricingDetailPerAdult{
						Availability: segment.PricingDetailPerAdult.Availability,
						FareBasis:    segment.PricingDetailPerAdult.FareBasis,
						FareClass:    segment.PricingDetailPerAdult.FareClass,
						TravelClass:  segment.PricingDetailPerAdult.TravelClass,
					},
				})
			}

			services = append(services, &pbType.Service{
				Segments: segments,
			})
		}

		off = append(off, &pbType.OfferItem{
			Services: services,
			Price: &pbType.Price{
				Total:      offers.Price.Total,
				TotalTaxes: offers.Price.TotalTaxes,
			},
			PricePerAdult: &pbType.Price{
				Total:      offers.PricePerAdult.Total,
				TotalTaxes: offers.PricePerAdult.TotalTaxes,
			},
		})
	}

	return &pbType.FlightLowFareSearchResult{
		Id:         resp.Id,
		OfferItems: off,
		Type:       resp.Type,
	}, nil
}

type TransportMiddleware func (pbFunc.AmadeusServiceServer) pbFunc.AmadeusServiceServer