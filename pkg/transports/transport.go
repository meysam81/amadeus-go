package transports

import (
	pbFunc "amadeus-go/api/amadeus/func"
	pbType "amadeus-go/api/amadeus/type"
	"amadeus-go/pkg/endpoints"
	sv "amadeus-go/pkg/services"

	"context"
	"os"

	"github.com/go-kit/kit/log"
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
	resp := response.(*sv.FlightLowFareSearchResponse)

	var datas []*pbType.Data
	for _, data := range resp.Data {

		var offerItems []*pbType.OfferItem
		for _, offers := range data.OfferItems {

			var services []*pbType.Service
			for _, service := range offers.Services {

				var segments []*pbType.Segment
				for _, segment := range service.Segments {

					segments = append(segments, &pbType.Segment{
						FlightSegment: &pbType.FlightSegment{
							Duration: segment.FlightSegment.Duration,
							Number:   segment.FlightSegment.Number,
							Aircraft: &pbType.Aircraft{
								Code: segment.FlightSegment.Aircraft.Code,
							},
							Arrival: &pbType.DepartureArrival{
								At:       segment.FlightSegment.Arrival.At,
								IataCode: segment.FlightSegment.Arrival.IataCode,
								Terminal: segment.FlightSegment.Arrival.Terminal,
							},
							Departure: &pbType.DepartureArrival{
								At:       segment.FlightSegment.Departure.At,
								IataCode: segment.FlightSegment.Departure.IataCode,
								Terminal: segment.FlightSegment.Departure.Terminal,
							},
							CarrierCode: segment.FlightSegment.CarrierCode,
							Operating: &pbType.Operating{
								CarrierCode: segment.FlightSegment.Operating.CarrierCode,
								Number:      segment.FlightSegment.Operating.Number,
							},
						},
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

			offerItems = append(offerItems, &pbType.OfferItem{
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

		datas = append(datas, &pbType.Data{
			Id:         data.Id,
			Type:       data.Type,
			OfferItems: offerItems,
		})
	}

	var dictionaries pbType.Dictionaries
	dictionaries.Aircrafts = make(map[string]string)
	dictionaries.Locations = make(map[string]*pbType.LocationDetail)
	dictionaries.Carriers = make(map[string]string)
	dictionaries.Currencies = make(map[string]string)
	for k, v := range resp.Dictionaries.Aircraft {
		dictionaries.Aircrafts[k] = v
	}
	for k, v := range resp.Dictionaries.Locations {
		if _, ok := dictionaries.Locations[k]; !ok {
			dictionaries.Locations[k] = &pbType.LocationDetail{
				Detail: make(map[string]string),
			}
		}
		for subK, subV := range v {
			dictionaries.Locations[k] = &pbType.LocationDetail{
				Detail: map[string]string{subK: subV},
			}
		}
	}
	for k, v := range resp.Dictionaries.Carriers {
		dictionaries.Aircrafts[k] = v
	}
	for k, v := range resp.Dictionaries.Currencies {
		dictionaries.Aircrafts[k] = v
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
		},
		Currency: resp.Meta.Currency,
		Defaults: &pbType.Defaults{
			Adults:  resp.Meta.Defaults.Adults,
			NonStop: resp.Meta.Defaults.NonStop,
		},
	}

	return &pbType.FlightLowFareSearchResult{
		Data:         datas,
		Dictionaries: &dictionaries,
		Meta:         &meta,
	}, nil

}

type TransportMiddleware func(pbFunc.AmadeusServiceServer) pbFunc.AmadeusServiceServer
