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
	"strings"
)

type grpcServer struct {
	FlightLowFareSearchHandler             grpcTransport.Handler
	FlightInspirationSearchHandler         grpcTransport.Handler
	FlightCheapestDateSearchHandler        grpcTransport.Handler
	FlightMostSearchedDestinationsHandler  grpcTransport.Handler
	FlightMostSearchedByDestinationHandler grpcTransport.Handler
	FlightCheckInLinksHandler              grpcTransport.Handler
	FlightMostTraveledDestinationsHandler  grpcTransport.Handler
	FlightMostBookedDestinationsHandler    grpcTransport.Handler
	FlightBusiestTravelingPeriodHandler    grpcTransport.Handler
	AirportNearestRelevantHandler          grpcTransport.Handler
	AirportAndCitySearchHandler            grpcTransport.Handler
	AirlineCodeLookupHandler               grpcTransport.Handler
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

func (s *grpcServer) FlightCheapestDateSearch(ctx context.Context, req *pbFunc.FlightCheapestDateSearchRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightCheapestDateSearchHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightMostSearchedDestinations(ctx context.Context, req *pbFunc.FlightMostSearchedDestinationsRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightMostSearchedDestinationsHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightMostSearchedByDestination(ctx context.Context, req *pbFunc.FlightMostSearchedByDestinationRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightMostSearchedByDestinationHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	response := resp.(*pbType.Response)
	return response, nil
}

func (s *grpcServer) FlightCheckInLinks(ctx context.Context, req *pbFunc.FlightCheckInLinksRequest) (*pbType.Response, error) {
	_, resp, err := s.FlightCheckInLinksHandler.ServeGRPC(ctx, req)
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

func (s *grpcServer) AirlineCodeLookup(ctx context.Context, req *pbFunc.AirlineCodeLookupRequest) (*pbType.Response, error) {
	_, resp, err := s.AirlineCodeLookupHandler.ServeGRPC(ctx, req)
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
			encodeResponse,
		),
		FlightInspirationSearchHandler: grpcTransport.NewServer(
			endpoints.FlightInspirationSearchEndpoint,
			decodeFlightInspirationSearchRequest,
			encodeResponse,
		),
		FlightCheapestDateSearchHandler: grpcTransport.NewServer(
			endpoints.FlightCheapestDateSearchEndpoint,
			decodeFlightCheapestDateSearchRequest,
			encodeResponse,
		),
		FlightMostTraveledDestinationsHandler: grpcTransport.NewServer(
			endpoints.FlightMostTraveledDestinationsEndpoint,
			decodeFlightMostTraveledDestinationsRequest,
			encodeResponse,
		),
		FlightMostSearchedDestinationsHandler: grpcTransport.NewServer(
			endpoints.FlightMostSearchedDestinationsEndpoint,
			decodeFlightMostSearchedDestinationsRequest,
			encodeResponse,
		),
		FlightMostSearchedByDestinationHandler: grpcTransport.NewServer(
			endpoints.FlightMostSearchedByDestinationEndpoint,
			decodeFlightMostSearchedByDestinationRequest,
			encodeResponse,
		),
		FlightCheckInLinksHandler: grpcTransport.NewServer(
			endpoints.FlightCheckInLinksEndpoint,
			decodeFlightCheckInLinksRequest,
			encodeResponse,
		),
		FlightMostBookedDestinationsHandler: grpcTransport.NewServer(
			endpoints.FlightMostBookedDestinationsEndpoint,
			decodeFlightMostBookedDestinationsRequest,
			encodeResponse,
		),
		FlightBusiestTravelingPeriodHandler: grpcTransport.NewServer(
			endpoints.FlightBusiestTravelingPeriodEndpoint,
			decodeFlightBusiestTravelingPeriodRequest,
			encodeResponse,
		),
		AirportNearestRelevantHandler: grpcTransport.NewServer(
			endpoints.AirportNearestRelevantEndpoint,
			decodeAirportNearestRelevantRequest,
			encodeResponse,
		),
		AirportAndCitySearchHandler: grpcTransport.NewServer(
			endpoints.AirportAndCitySearchEndpoint,
			decodeAirportAndCitySearchRequest,
			encodeResponse,
		),
		AirlineCodeLookupHandler: grpcTransport.NewServer(
			endpoints.AirlineCodeLookupEndpoint,
			decodeAirlineCodeLookupRequest,
			encodeResponse,
		),
	}

	return
}

// =============================================================================

// watch out for this method, it is long and hard to understand cause we are
// refactoring the response from the server to that of a protobuf
func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp, ok := response.(*sv.Response)
	if !ok {
		return nil, errors.New("couldn't convert response to <Response>")
	}

	// ******************** Step 1: Data ********************
	var datas []*pbType.Data
	for _, data := range resp.Data {

		var offerItems []*pbType.OfferItem
		if data.OfferItems != nil {
			for _, offers := range data.OfferItems {

				var services []*pbType.Service
				if offers.Services != nil {
					for _, service := range offers.Services {

						var segments []*pbType.Segment
						if service.Segments != nil {
							for _, segment := range service.Segments {

								var pricingDetailPerAdult pbType.PricingDetailPerAdult
								if segment.PricingDetailPerAdult != nil {
									pricingDetailPerAdult = pbType.PricingDetailPerAdult{
										Availability: segment.PricingDetailPerAdult.Availability,
										FareBasis:    segment.PricingDetailPerAdult.FareBasis,
										FareClass:    segment.PricingDetailPerAdult.FareClass,
										TravelClass:  segment.PricingDetailPerAdult.TravelClass,
									}
								}

								var flightSegment pbType.FlightSegment
								if segment.FlightSegment != nil {
									flightSegment = pbType.FlightSegment{
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
									}
								}

								segments = append(segments, &pbType.Segment{
									FlightSegment:         &flightSegment,
									PricingDetailPerAdult: &pricingDetailPerAdult,
								})
							} // endfor service.Segments
						} // endif service.Segments

						services = append(services, &pbType.Service{
							Segments: segments,
						})
					} // endfor offers.Services
				} // endif offers.Services

				var price, pricePerAdult pbType.Price
				if offers.Price != nil {
					price = pbType.Price{
						Total:      offers.Price.Total,
						TotalTaxes: offers.Price.TotalTaxes,
					}
				}
				if offers.PricePerAdult != nil {
					pricePerAdult = pbType.Price{
						Total:      offers.PricePerAdult.Total,
						TotalTaxes: offers.PricePerAdult.TotalTaxes,
					}
				}

				offerItems = append(offerItems, &pbType.OfferItem{
					Services:      services,
					Price:         &price,
					PricePerAdult: &pricePerAdult,
				})
			} // endfor data.OfferItems
		} // endif data.OfferItems

		var geoCode pbType.GeoCode
		if data.GeoCode != nil {
			geoCode = pbType.GeoCode{
				Latitude:  data.GeoCode.Latitude,
				Longitude: data.GeoCode.Longitude,
			}
		}

		var distance pbType.Distance
		if data.Distance != nil {
			distance = pbType.Distance{
				Unit:  data.Distance.Unit,
				Value: data.Distance.Value,
			}
		}

		var price pbType.Price
		if data.Price != nil {
			price = pbType.Price{
				Total:      data.Price.Total,
				TotalTaxes: data.Price.TotalTaxes,
			}
		}

		var links pbType.Links
		if data.Links != nil {
			links = pbType.Links{
				Self:               data.Links.Self,
				Next:               data.Links.Next,
				Last:               data.Links.Last,
				FlightDates:        data.Links.FlightDates,
				FlightOffers:       data.Links.FlightOffers,
				FlightDestinations: data.Links.FlightDestinations,
			}
		}

		var self pbType.Self
		if data.Self != nil {
			self.Href = data.Self.Href
			for _, m := range data.Self.Methods {
				self.Methods = append(self.Methods, m)
			}
		}

		var address pbType.Address
		if data.Address != nil {
			address = pbType.Address{
				CityName:    data.Address.CityName,
				CityCode:    data.Address.CityCode,
				CountryName: data.Address.CountryName,
				CountryCode: data.Address.CountryCode,
				StateCode:   data.Address.StateCode,
				RegionCode:  data.Address.RegionCode,
			}
		}

		var analytics pbType.Analytics
		if data.Analytics != nil {
			if data.Analytics.Flights != nil {
				analytics.Flights.Score = data.Analytics.Flights.Score
			}

			if data.Analytics.Travelers != nil {
				analytics.Travelers.Score = data.Analytics.Travelers.Score
			}

			if data.Analytics.Searches != nil {
				analytics.Searches.Score = data.Analytics.Searches.Score
				if data.Analytics.Searches.NumberOfSearches != nil {

					perTripDuration := make(map[string]string)
					perDaysInAdvance := make(map[string]string)
					for k, v := range data.Analytics.Searches.NumberOfSearches.PerTripDuration {
						perTripDuration[k] = v
					}
					for k, v := range data.Analytics.Searches.NumberOfSearches.PerDaysInAdvance {
						perDaysInAdvance[k] = v
					}

					analytics.Searches.NumberOfSearches = &pbType.NumberOfSearches{
						PerDaysInAdvance: perDaysInAdvance,
						PerTripDuration:  perTripDuration,
					}
				} // endif data.Analytics.Searches.NumberOfSearches != nil
			} // endif data.Analytics.Searches != nil
		} // endif data.Analytics != nil

		params := make(map[string]*pbType.ParamDetail)
		for k, v := range data.Parameters {
			p := pbType.ParamDetail{
				Type:        v.Type,
				Description: v.Description,
				Format:      v.Format,
			}
			params[k] = &p
		}

		newData := pbType.Data{
			Id:             data.Id,
			Type:           data.Type,
			OfferItems:     offerItems,
			Destination:    data.Destination,
			SubType:        data.SubType,
			Analytics:      &analytics,
			Period:         data.Period,
			Name:           data.Name,
			DetailedName:   data.DetailedName,
			TimeZoneOffset: data.TimeZoneOffset,
			IataCode:       data.IataCode,
			GeoCode:        &geoCode,
			Address:        &address,
			Distance:       &distance,
			Relevance:      data.Relevance,
			Origin:         data.Origin,
			DepartureDate:  data.DepartureDate,
			ReturnDate:     data.ReturnDate,
			Price:          &price,
			Links:          &links,
			Self:           &self,
			Href:           data.Href,
			Channel:        data.Channel,
			Parameters:     params,
		}
		datas = append(datas, &newData)
	} // endfor resp.Data

	// ******************** Step 2: Dictionary ********************
	var dictionaries pbType.Dictionaries
	if resp.Dictionaries != nil {
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
	} // endif resp.Dictionaries != nil

	// ******************** Step 3: Meta ********************
	var meta pbType.Meta
	if resp.Meta != nil {
		var links pbType.Links
		if resp.Meta.Links != nil {
			links = pbType.Links{
				Self:               resp.Meta.Links.Self,
				Next:               resp.Meta.Links.Next,
				Last:               resp.Meta.Links.Last,
				FlightDates:        resp.Meta.Links.FlightDates,
				FlightOffers:       resp.Meta.Links.FlightOffers,
				FlightDestinations: resp.Meta.Links.FlightDestinations,
			}
		}

		var defaults pbType.Defaults
		if resp.Meta.Defaults != nil {
			defaults = pbType.Defaults{
				Adults:  resp.Meta.Defaults.Adults,
				NonStop: resp.Meta.Defaults.NonStop,
			}
		}

		meta = pbType.Meta{
			Links:    &links,
			Currency: resp.Meta.Currency,
			Defaults: &defaults,
			Count:    resp.Meta.Count,
		}
	}

	// ******************** Step 4: Warnings and Erros ********************
	var errs, warnings []*pbType.ErrorWarning
	if resp.Warnings != nil {
		for _, w := range resp.Warnings {
			warning := pbType.ErrorWarning{
				Title:  w.Title,
				Status: w.Status,
				Code:   w.Code,
				Detail: w.Detail,
			}
			if w.Source != nil {
				warning.Source = &pbType.Source{
					Example:   w.Source.Example,
					Parameter: w.Source.Parameter,
					Pointer:   w.Source.Pointer,
				}
			}
			warnings = append(warnings, &warning)
		} // endfor
	}
	if resp.Errors != nil {
		for _, w := range resp.Errors {
			err := pbType.ErrorWarning{
				Title:  w.Title,
				Status: w.Status,
				Code:   w.Code,
				Detail: w.Detail,
			}
			if w.Source != nil {
				err.Source = &pbType.Source{
					Example:   w.Source.Example,
					Parameter: w.Source.Parameter,
					Pointer:   w.Source.Pointer,
				}
			}
			errs = append(errs, &err)
		} // endfor
	}

	return &pbType.Response{
		Data:         datas,
		Dictionaries: &dictionaries,
		Meta:         &meta,
		Errors:       errs,
		Warnings:     warnings,
	}, nil
}

func decodeFlightLowFareSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightLowFareSearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightLowFareSearchRequest>")
	}

	// check required fields
	if emptyString(req.Origin) {
		return nil, errors.New("origin is a required field")
	}
	if emptyString(req.Destination) {
		return nil, errors.New("destination is a required field")
	}
	if emptyString(req.DepartureDate) {
		return nil, errors.New("departureDate is a required field")
	}

	request := sv.FlightLowFareSearchRequest{
		Origin:        req.Origin,
		Destination:   req.Destination,
		DepartureDate: req.DepartureDate,
	}

	// check optional fields
	if !emptyString(req.ReturnDate) {
		request.ReturnDate = req.ReturnDate
	}
	if !emptyString(req.ArrivalBy) {
		request.ArrivalBy = req.ArrivalBy
	}
	if !emptyString(req.ReturnBy) {
		request.ReturnBy = req.ReturnBy
	}
	if req.Adults > 0 {
		request.Adults = req.Adults
	}
	if req.Children > 0 {
		request.Children = req.Children
	}
	if req.Infants > 0 {
		request.Infants = req.Infants
	}
	if req.Seniors > 0 {
		request.Seniors = req.Seniors
	}
	switch req.TravelClass {
	case pbType.TravelClass_ECONOMY:
		*request.TravelClass = sv.TravelClass_ECONOMY
	case pbType.TravelClass_PREMIUM_ECONOMY:
		*request.TravelClass = sv.TravelClass_PREMIUM_ECONOMY
	case pbType.TravelClass_BUSINESS:
		*request.TravelClass = sv.TravelClass_BUSINESS
	case pbType.TravelClass_FIRST:
		*request.TravelClass = sv.TravelClass_FIRST
	default:
	}
	if !emptyString(req.IncludeAirlines) {
		request.IncludeAirlines = req.IncludeAirlines
	}
	if !emptyString(req.ExcludeAirlines) {
		request.ExcludeAirlines = req.ExcludeAirlines
	}
	if req.NonStop {
		request.NonStop = req.NonStop
	}
	if !emptyString(req.Currency) {
		request.Currency = req.Currency
	}
	if req.MaxPrice > 0 {
		request.MaxPrice = req.MaxPrice
	}
	if req.Max > 0 && req.Max <= 250 {
		request.Max = req.Max
	}

	return &request, nil
}

func decodeFlightInspirationSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightInspirationSearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightInspirationSearchRequest>")
	}

	if emptyString(req.Origin) {
		return nil, errors.New("origin is a required field")
	}

	request := sv.FlightInspirationSearchRequest{
		Origin: req.Origin,
	}

	if !emptyString(req.DepartureDate) {
		request.DepartureDate = req.DepartureDate
	}
	if req.OneWay {
		request.OneWay = req.OneWay
	}
	if !emptyString(req.Duration) {
		request.Duration = req.Duration
	}
	if req.NonStop {
		request.NonStop = req.NonStop
	}
	if req.MaxPrice > 0 {
		request.MaxPrice = req.MaxPrice
	}
	if !emptyString(req.Currency) {
		request.Currency = req.Currency
	}

	return &request, nil
}

func decodeFlightCheapestDateSearchRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightCheapestDateSearchRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightCheapestDateSearchRequest>")
	}

	if emptyString(req.Origin) {
		return nil, errors.New("origin is a required field")
	}
	if emptyString(req.Destination) {
		return nil, errors.New("destination is a required field")
	}

	request := sv.FlightCheapestDateSearchRequest{
		Origin:      req.Origin,
		Destination: req.Destination,
	}

	if !emptyString(req.DepartureDate) {
		request.DepartureDate = req.DepartureDate
	}
	if req.OneWay {
		request.OneWay = req.OneWay
	}
	if !emptyString(req.Duration) {
		request.Duration = req.Duration
	}
	if req.NonStop {
		request.NonStop = req.NonStop
	}
	if req.MaxPrice > 0 {
		request.MaxPrice = req.MaxPrice
	}
	if !emptyString(req.Currency) {
		request.Currency = req.Currency
	}
	switch req.ViewBy {
	case pbType.ViewBy_DATE:
		*request.ViewBy = sv.ViewBy_DATE
	case pbType.ViewBy_DURATION:
		*request.ViewBy = sv.ViewBy_DURATION
	case pbType.ViewBy_WEEK:
		*request.ViewBy = sv.ViewBy_WEEK
	default:
	}

	return &request, nil
}

func decodeFlightMostSearchedDestinationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostSearchedDestinationsRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostSearchedDestinationsRequest>")
	}

	if emptyString(req.OriginCityCode) {
		return nil, errors.New("originCityCode is a required field")
	}
	if emptyString(req.SearchPeriod) {
		return nil, errors.New("searchPeriod is a required field")
	}
	if emptyString(req.MarketCountryCode) {
		return nil, errors.New("marketCountryCode is a required field")
	}

	request := sv.FlightMostSearchedDestinationsRequest{
		OriginCityCode:    req.OriginCityCode,
		SearchPeriod:      req.SearchPeriod,
		MarketCountryCode: req.MarketCountryCode,
	}

	if req.Max > 0 {
		request.Max = req.Max
	}
	if !emptyString(req.Fields) {
		request.Fields = req.Fields
	}
	if request.PageLimit > 0 {
		request.PageLimit = req.PageLimit
	}
	if request.PageOffset > 0 {
		request.PageOffset = req.PageOffset
	}

	return &request, nil
}

func decodeFlightMostSearchedByDestinationRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostSearchedByDestinationRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostSearchedByDestination>")
	}

	if emptyString(req.OriginCityCode) {
		return nil, errors.New("originCityCode is a required field")
	}
	if emptyString(req.DestinationCityCode) {
		return nil, errors.New("destinationCityCode is a required field")
	}
	if emptyString(req.SearchPeriod) {
		return nil, errors.New("searchPeriod is a required field")
	}
	if emptyString(req.MarketCountryCode) {
		return nil, errors.New("marketCountryCode is a required field")
	}

	request := sv.FlightMostSearchedByDestinationRequest{
		OriginCityCode:      req.OriginCityCode,
		DestinationCityCode: req.DestinationCityCode,
		SearchPeriod:        req.SearchPeriod,
		MarketCountryCode:   req.MarketCountryCode,
	}

	if !emptyString(req.Fields) {
		request.Fields = req.Fields
	}

	return &sv.FlightMostSearchedByDestinationRequest{
		MarketCountryCode:   req.MarketCountryCode,
		SearchPeriod:        req.SearchPeriod,
		OriginCityCode:      req.OriginCityCode,
		DestinationCityCode: req.DestinationCityCode,
	}, nil
}

func decodeFlightMostTraveledDestinationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostTraveledDestinationsRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostTraveledDestinationsRequest>")
	}

	if emptyString(req.OriginCityCode) {
		return nil, errors.New("originCityCode is a required field")
	}
	if emptyString(req.Period) {
		return nil, errors.New("period is a required field")
	}

	request := sv.FlightMostTraveledDestinationsRequest{
		OriginCityCode: req.OriginCityCode,
		Period:         req.Period,
	}

	if req.Max > 0 {
		request.Max = req.Max
	}
	if !emptyString(req.Fields) {
		request.Fields = req.Fields
	}
	if req.PageLimit > 0 {
		request.PageLimit = req.PageLimit
	}
	if req.PageOffset > 0 {
		request.PageOffset = req.PageOffset
	}
	switch req.Sort {
	case pbType.Sort_FLIGHTS:
		*request.Sort = sv.Sort_FLIGHTS
	case pbType.Sort_TRAVELERS:
		*request.Sort = sv.Sort_TRAVELERS
	default:
	}

	return &request, nil
}

func decodeFlightMostBookedDestinationsRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightMostBookedDestinationsRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightMostBookedDestinationsRequest>")
	}

	if emptyString(req.OriginCityCode) {
		return nil, errors.New("originCityCode is a required field")
	}
	if emptyString(req.Period) {
		return nil, errors.New("period is a required field")
	}

	request := sv.FlightMostBookedDestinationsRequest{
		OriginCityCode: req.OriginCityCode,
		Period:         req.Period,
	}

	if req.Max > 0 {
		request.Max = req.Max
	}
	if !emptyString(req.Fields) {
		request.Fields = req.Fields
	}
	if req.PageLimit > 0 {
		request.PageLimit = req.PageLimit
	}
	if req.PageOffset > 0 {
		request.PageOffset = req.PageOffset
	}
	switch req.Sort {
	case pbType.Sort_FLIGHTS:
		*request.Sort = sv.Sort_FLIGHTS
	case pbType.Sort_TRAVELERS:
		*request.Sort = sv.Sort_TRAVELERS
	default:
	}

	return &request, nil
}

func decodeFlightBusiestTravelingPeriodRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightBusiestTravelingPeriodRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightBusiestTravelingPeriodRequest>")
	}

	if emptyString(req.CityCode) {
		return nil, errors.New("cityCode is a required field")
	}
	if emptyString(req.Period) {
		return nil, errors.New("period is a required field")
	}

	request := sv.FlightBusiestTravelingPeriodRequest{
		CityCode: req.CityCode,
		Period:   req.Period,
	}

	switch req.Direction {
	case pbType.Direction_ARRIVING:
		*request.Direction = sv.Direction_ARRIVING
	case pbType.Direction_DEPARTING:
		*request.Direction = sv.Direction_DEPARTING
	default:
	}

	return &request, nil
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

func decodeFlightCheckInLinksRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.FlightCheckInLinksRequest)
	if !ok {
		return nil, errors.New("your request is not of type <FlightCheckInLinksRequest>")
	}

	if emptyString(req.AirlineCode) {
		return nil, errors.New("airlineCode is a required field")
	}

	request := sv.FlightCheckInLinksRequest{
		AirlineCode: req.AirlineCode,
	}

	if !emptyString(req.Language) {
		request.Language = req.Language
	}

	return &request, nil
}

func decodeAirlineCodeLookupRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req, ok := grpcReq.(*pbFunc.AirlineCodeLookupRequest)
	if !ok {
		return nil, errors.New("your request is not of type <AirlineCodeLookupRequest>")
	}
	return &sv.AirlineCodeLookupRequest{
		AirlineCodes: req.AirlineCodes,
	}, nil
}

func emptyString(s string) bool {
	return strings.Compare(s, "") == 0
}
