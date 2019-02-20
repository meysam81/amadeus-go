package transports

import (
	pbFunc "amadeus-go/api/amadeus/func"
	pbType "amadeus-go/api/amadeus/type"
	"amadeus-go/pkg/endpoints"
	srv "amadeus-go/pkg/services"

	"context"

	"github.com/go-kit/kit/endpoint"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) srv.AmadeusService {
	var flightLowFareSearchEndpoint endpoint.Endpoint
	{
		flightLowFareSearchEndpoint = grpcTransport.NewClient(
			conn,
			"amadeus.func.AmadeusService",
			"FlightLowFareSearch",
			encodeRequest,
			decodeResponse,
			pbType.FlightLowFareSearchResult{},
		).Endpoint()
	}

	return endpoints.AmadeusEndpointSet{FlightLowFareSearchEndpoint: flightLowFareSearchEndpoint}
}

func encodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(*srv.FlightLowFareSearchRequest)
	return &pbFunc.FlightLowFareSearchRequest{
		Origin:        req.Origin,
		DepartureDate: req.DepartureDate,
		Destination:   req.Destination,
		ReturnDate:    req.ReturnDate,
	}, nil
}

func decodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pbType.FlightLowFareSearchResult)

	var datas []*srv.Data
	for _, data := range resp.Data {

		var offerItems []*srv.OfferItem
		for _, offers := range data.OfferItems {

			var services []*srv.Service
			for _, service := range offers.Services {

				var segments []*srv.Segment
				for _, segment := range service.Segments {

					segments = append(segments, &srv.Segment{
						FlightSegment: &srv.FlightSegment{
							Duration: segment.FlightSegment.Duration,
							Number:   segment.FlightSegment.Number,
							Aircraft: &srv.Aircraft{
								Code: segment.FlightSegment.Aircraft.Code,
							},
							Arrival: &srv.DepartureArrival{
								At:       segment.FlightSegment.Arrival.At,
								IataCode: segment.FlightSegment.Arrival.IataCode,
								Terminal: segment.FlightSegment.Arrival.Terminal,
							},
							Departure: &srv.DepartureArrival{
								At:       segment.FlightSegment.Departure.At,
								IataCode: segment.FlightSegment.Departure.IataCode,
								Terminal: segment.FlightSegment.Departure.Terminal,
							},
							CarrierCode: segment.FlightSegment.CarrierCode,
							Operating: &srv.Operating{
								CarrierCode: segment.FlightSegment.Operating.CarrierCode,
								Number:      segment.FlightSegment.Operating.Number,
							},
						},
						PricingDetailPerAdult: &srv.PricingDetailPerAdult{
							Availability: segment.PricingDetailPerAdult.Availability,
							FareBasis:    segment.PricingDetailPerAdult.FareBasis,
							FareClass:    segment.PricingDetailPerAdult.FareClass,
							TravelClass:  segment.PricingDetailPerAdult.TravelClass,
						},
					})
				}

				services = append(services, &srv.Service{
					Segments: segments,
				})
			}

			offerItems = append(offerItems, &srv.OfferItem{
				Services: services,
				Price: &srv.Price{
					Total:      offers.Price.Total,
					TotalTaxes: offers.Price.TotalTaxes,
				},
				PricePerAdult: &srv.Price{
					Total:      offers.PricePerAdult.Total,
					TotalTaxes: offers.PricePerAdult.TotalTaxes,
				},
			})
		}

		datas = append(datas, &srv.Data{
			Id:         data.Id,
			Type:       data.Type,
			OfferItems: offerItems,
		})
	}

	var dictionaries srv.Dictionaries
	_ = mapstructure.Decode(resp.Dictionaries, &dictionaries)

	meta := srv.Meta{
		Links: &srv.Links{
			Self: resp.Meta.Links.Self,
		},
		Currency: resp.Meta.Currency,
		Defaults: &srv.Defaults{
			Adults:  resp.Meta.Defaults.Adults,
			NonStop: resp.Meta.Defaults.NonStop,
		},
	}

	return &srv.FlightLowFareSearchResponse{
		Data:         datas,
		Dictionaries: &dictionaries,
		Meta:         &meta,
	}, nil
}
