package transports

import (
	pbFunc "amadeus-go/api/amadeus/func"
	pbType "amadeus-go/api/amadeus/type"
	"amadeus-go/pkg/endpoints"
	sv "amadeus-go/pkg/services"
	"context"
	"errors"
	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	FlightLowFareSearchHandler            grpcTransport.Handler
	FlightInspirationSearchHandler        grpcTransport.Handler
	FlightMostTraveledDestinationsHandler grpcTransport.Handler
	FlightMostBookedDestinationsHandler   grpcTransport.Handler
	FlightBusiestTravelingPeriodHandler   grpcTransport.Handler
	AirportNearestRelevantHandler         grpcTransport.Handler
	AirportAndCitySearchHandler           grpcTransport.Handler
}

func (s *grpcServer) FlightLowFareSearch(ctx context.Context, req *pbFunc.FlightLowFareSearchRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightLowFareSearchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightInspirationSearch(ctx context.Context, req *pbFunc.FlightInspirationSearchRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightLowFareSearchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightMostTraveledDestinations(ctx context.Context, req *pbFunc.FlightMostTraveledDestinationsRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightLowFareSearchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightMostBookedDestinations(ctx context.Context, req *pbFunc.FlightMostBookedDestinationsRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightMostBookedDestinationsHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightBusiestTravelingPeriod(ctx context.Context, req *pbFunc.FlightBusiestTravelingPeriodRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightBusiestTravelingPeriodHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) AirportNearestRelevant(ctx context.Context, req *pbFunc.AirportNearestRelevantRequest) (*pbType.Response, error) {
	_, resp, err := s.AirportNearestRelevantHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) AirportAndCitySearch(ctx context.Context, req *pbFunc.AirportAndCitySearchRequest) (*pbType.Response, error) {
	_, resp, err := s.AirportNearestRelevantHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func NewGRPCServer(endpoints *endpoints.AmadeusEndpointSet, logger log.Logger) (s pbFunc.AmadeusServiceServer) {
	s = &grpcServer{
		FlightLowFareSearchHandler: grpcTransport.NewServer(
			endpoints.FlightLowFareSearchEndpoint,
			decodeFlightLowFareSearchRequest,
			encodeFlightLowFareSearchResponse,
		),
		FlightInspirationSearchHandler: grpcTransport.NewServer(
			endpoints.FlightInspirationSearchEndpoint,
			decodeFlightInspirationSearchRequest,
			encodeFlightInspirationSearchResponse,
		),
		FlightMostTraveledDestinationsHandler: grpcTransport.NewServer(
			endpoints.FlightMostTraveledDestinationsEndpoint,
			decodeFlightMostTraveledDestinationsRequest,
			encodeFlightMostTraveledDestinationsResponse,
		),
		FlightMostBookedDestinationsHandler: grpcTransport.NewServer(
			endpoints.FlightMostBookedDestinationsEndpoint,
			decodeFlightMostBookedDestinationsRequest,
			encodeFlightMostBookedDestinationsResponse,
		),
		FlightBusiestTravelingPeriodHandler: grpcTransport.NewServer(
			endpoints.FlightBusiestTravelingPeriodEndpoint,
			decodeFlightBusiestTravelingPeriodRequest,
			encodeFlightBusiestTravelingPeriodResponse,
		),
		AirportNearestRelevantHandler: grpcTransport.NewServer(
			endpoints.AirportNearestRelevantEndpoint,
			decodeAirportNearestRelevantRequest,
			encodeAirportNearestRelevantResponse,
		),
		AirportAndCitySearchHandler: grpcTransport.NewServer(
			endpoints.AirportAndCitySearchEndpoint,
			decodeAirportAndCitySearchRequest,
			encodeAirportAndCitySearchResponse,
		),
	}

	return
}

// =============================================================================
func decodeFlightLowFareSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightLowFareSearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightLowFareSearchRequest>")
	}
	return &sv.FlightLowFareSearchRequest{
		Origin:        req.Origin,
		DepartureDate: req.DepartureDate,
		Destination:   req.Destination,
		ReturnDate:    req.ReturnDate,
	}, nil
}

func encodeFlightLowFareSearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

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
	for k, v := range resp.Dictionaries.Aircrafts {
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

	return &pbType.Response{
		Data:         datas,
		Dictionaries: &dictionaries,
		Meta:         &meta,
	}, nil
}

func decodeFlightInspirationSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightInspirationSearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightInspirationSearchRequest>")
	}
	return &sv.FlightInspirationSearchRequest{
		Origin:   req.Origin,
		MaxPrice: req.MaxPrice,
	}, nil
}

func encodeFlightInspirationSearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {
		datas = append(datas, &pbType.Data{
			Type:          data.Type,
			Origin:        data.Origin,
			Destination:   data.Destination,
			DepartureDate: data.DepartureDate,
			ReturnDate:    data.ReturnDate,
			Price: &pbType.Price{
				Total: data.Price.Total,
			},
			Links: &pbType.Links{
				FlightDates:  data.Links.FlightDates,
				FlightOffers: data.Links.FlightOffers,
			},
		})
	}

	var dictionaries pbType.Dictionaries
	dictionaries.Locations = make(map[string]*pbType.LocationDetail)
	dictionaries.Currencies = make(map[string]string)
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
	for k, v := range resp.Dictionaries.Currencies {
		dictionaries.Aircrafts[k] = v
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
		},
		Currency: resp.Meta.Currency,
		Defaults: &pbType.Defaults{
			DepartureDate: resp.Meta.Defaults.DepartureDate,
			OneWay:        resp.Meta.Defaults.OneWay,
			Duration:      resp.Meta.Defaults.Duration,
			NonStop:       resp.Meta.Defaults.NonStop,
			ViewBy:        resp.Meta.Defaults.ViewBy,
		},
	}

	return &pbType.Response{
		Data:         datas,
		Dictionaries: &dictionaries,
		Meta:         &meta,
	}, nil
}

func decodeFlightMostTraveledDestinationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostTraveledDestinationsRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostTraveledDestinationsRequest>")
	}
	return &sv.FlightMostTraveledDestinationsRequest{
		OriginCityCode: req.OriginCityCode,
		Period:         req.Period,
	}, nil
}

func encodeFlightMostTraveledDestinationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {
		datas = append(datas, &pbType.Data{
			Type:        data.Type,
			Destination: data.Destination,
			SubType:     data.SubType,
			Analytics: &pbType.Analytics{
				Flights: &pbType.Score{
					Score: data.Analytics.Flights.Score,
				},
				Travelers: &pbType.Score{
					Score: data.Analytics.Travelers.Score,
				},
			},
		})
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
		},
		Count: resp.Meta.Count,
	}

	return &pbType.Response{
		Data: datas,
		Meta: &meta,
	}, nil
}

func decodeFlightMostBookedDestinationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostBookedDestinationsRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostBookedDestinationsRequest>")
	}
	return &sv.FlightMostBookedDestinationsRequest{
		OriginCityCode: req.OriginCityCode,
		Period:         req.Period,
	}, nil
}

func encodeFlightMostBookedDestinationsResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {
		datas = append(datas, &pbType.Data{
			Type:        data.Type,
			Destination: data.Destination,
			SubType:     data.SubType,
			Analytics: &pbType.Analytics{
				Flights: &pbType.Score{
					Score: data.Analytics.Flights.Score,
				},
				Travelers: &pbType.Score{
					Score: data.Analytics.Travelers.Score,
				},
			},
		})
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
		},
		Count: resp.Meta.Count,
	}

	return &pbType.Response{
		Data: datas,
		Meta: &meta,
	}, nil
}

func decodeFlightBusiestTravelingPeriodRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightBusiestTravelingPeriodRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightBusiestTravelingPeriodRequest>")
	}
	return &sv.FlightBusiestTravelingPeriodRequest{
		CityCode:  req.CityCode,
		Period:    req.Period,
		Direction: req.Direction,
	}, nil
}

func encodeFlightBusiestTravelingPeriodResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {
		datas = append(datas, &pbType.Data{
			Type:   data.Type,
			Period: data.Period,
			Analytics: &pbType.Analytics{
				Flights: &pbType.Score{
					Score: data.Analytics.Flights.Score,
				},
				Travelers: &pbType.Score{
					Score: data.Analytics.Travelers.Score,
				},
			},
		})
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
		},
		Count: resp.Meta.Count,
	}

	return &pbType.Response{
		Data: datas,
		Meta: &meta,
	}, nil
}

func decodeAirportNearestRelevantRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.AirportNearestRelevantRequest)
	if !ok {
		return nil, errors.New("your request is not of type <AirportNearestRelevantRequest>")
	}
	return &sv.AirportNearestRelevantRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		Sort:      req.Sort,
	}, nil
}

func encodeAirportNearestRelevantResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {
		datas = append(datas, &pbType.Data{
			Type:           data.Type,
			SubType:        data.SubType,
			Name:           data.Name,
			DetailedName:   data.DetailedName,
			TimeZoneOffset: data.TimeZoneOffset,
			IataCode:       data.IataCode,
			GeoCode: &pbType.GeoCode{
				Latitude:  data.GeoCode.Latitude,
				Longitude: data.GeoCode.Longitude,
			},
			Address: &pbType.Address{
				CityName:    data.Address.CityName,
				CityCode:    data.Address.CityCode,
				CountryName: data.Address.CountryName,
				CountryCode: data.Address.CountryCode,
				RegionCode:  data.Address.RegionCode,
			},
			Distance: &pbType.Distance{
				Value: data.Distance.Value,
				Unit:  data.Distance.Unit,
			},
			Analytics: &pbType.Analytics{
				Flights: &pbType.Score{
					Score: data.Analytics.Flights.Score,
				},
				Travelers: &pbType.Score{
					Score: data.Analytics.Travelers.Score,
				},
			},
			Relevance: data.Relevance,
		})
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
			Next: resp.Meta.Links.Next,
			Last: resp.Meta.Links.Last,
		},
		Count: resp.Meta.Count,
	}

	return &pbType.Response{
		Data: datas,
		Meta: &meta,
	}, nil
}

func decodeAirportAndCitySearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.AirportAndCitySearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <AirportAndCitySearchRequest>")
	}
	return &sv.AirportAndCitySearchRequest{
		Keyword:     req.Keyword,
		SubType:     req.SubType,
		CountryCode: req.CountryCode,
	}, nil
}

func encodeAirportAndCitySearchResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	var datas []*pbType.Data
	for _, data := range resp.Data {

		var methods []string
		for _, method := range data.Self.Methods {
			methods = append(methods, method)
		}

		datas = append(datas, &pbType.Data{
			Type:           data.Type,
			SubType:        data.SubType,
			Name:           data.Name,
			DetailedName:   data.DetailedName,
			Id: data.Id,
			Self: &pbType.Self{
				Href: data.Self.Href,
				Methods: methods,
			},
			TimeZoneOffset: data.TimeZoneOffset,
			IataCode:       data.IataCode,
			GeoCode: &pbType.GeoCode{
				Latitude:  data.GeoCode.Latitude,
				Longitude: data.GeoCode.Longitude,
			},
			Address: &pbType.Address{
				CityName:    data.Address.CityName,
				CityCode:    data.Address.CityCode,
				CountryName: data.Address.CountryName,
				CountryCode: data.Address.CountryCode,
				StateCode:   data.Address.StateCode,
				RegionCode:  data.Address.RegionCode,
			},
			Analytics: &pbType.Analytics{
				Travelers: &pbType.Score{
					Score: data.Analytics.Travelers.Score,
				},
			},
		})
	}

	meta := pbType.Meta{
		Links: &pbType.Links{
			Self: resp.Meta.Links.Self,
			Next: resp.Meta.Links.Next,
			Last: resp.Meta.Links.Last,
		},
		Count: resp.Meta.Count,
	}

	return &pbType.Response{
		Data: datas,
		Meta: &meta,
	}, nil
}
