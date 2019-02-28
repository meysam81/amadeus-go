package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
)

type AmadeusService interface {
	FlightLowFareSearch(context.Context, *FlightLowFareSearchRequest) (*Response, error)
	FlightInspirationSearch(context.Context, *FlightInspirationSearchRequest) (*Response, error)
	FlightCheapestDateSearch(context.Context, *FlightCheapestDateSearchRequest) (*Response, error)
	FlightMostSearchedDestinations(context.Context, *FlightMostSearchedDestinationsRequest) (*Response, error)
	FlightMostSearchedByDestination(context.Context, *FlightMostSearchedByDestinationRequest) (*Response, error)
	FlightCheckInLinks(context.Context, *FlightCheckInLinksRequest) (*Response, error)
	FlightMostTraveledDestinations(context.Context, *FlightMostTraveledDestinationsRequest) (*Response, error)
	FlightMostBookedDestinations(context.Context, *FlightMostBookedDestinationsRequest) (*Response, error)
	FlightBusiestTravelingPeriod(context.Context, *FlightBusiestTravelingPeriodRequest) (*Response, error)
	AirportNearestRelevant(context.Context, *AirportNearestRelevantRequest) (*Response, error)
	AirportAndCitySearch(context.Context, *AirportAndCitySearchRequest) (*Response, error)
	AirlineCodeLookup(context.Context, *AirlineCodeLookupRequest) (*Response, error)
}

func (aSrv amadeusService) FlightLowFareSearch(_ context.Context, request *FlightLowFareSearchRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightLowFareSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("origin", request.Origin)
	q.Add("destination", request.Destination)
	q.Add("departureDate", request.DepartureDate)
	q.Add("returnDate", request.ReturnDate)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightInspirationSearch(_ context.Context, request *FlightInspirationSearchRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightInspirationSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("origin", request.Origin)
	q.Add("maxPrice", string(request.MaxPrice))
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightCheapestDateSearch(_ context.Context, request *FlightCheapestDateSearchRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightCheapestDateSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("origin", request.Origin)
	q.Add("destination", string(request.Destination))
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightMostSearchedDestinations(_ context.Context, request *FlightMostSearchedDestinationsRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostSearchedDestinations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("originCityCode", request.OriginCityCode)
	q.Add("searchPeriod", request.SearchPeriod)
	q.Add("marketCountryCode", request.MarketCountryCode)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightMostSearchedByDestination(_ context.Context, request *FlightMostSearchedByDestinationRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostSearchedByDestination)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("originCityCode", request.OriginCityCode)
	q.Add("destinationCityCode", request.DestinationCityCode)
	q.Add("searchPeriod", request.SearchPeriod)
	q.Add("marketCountryCode", request.MarketCountryCode)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightCheckInLinks(_ context.Context, request *FlightCheckInLinksRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightCheckInLists)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("airlineCode", request.AirlineCode)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightMostTraveledDestinations(_ context.Context, request *FlightMostTraveledDestinationsRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostTraveledDestinations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("originCityCode", request.OriginCityCode)
	q.Add("period", string(request.Period))
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightMostBookedDestinations(_ context.Context, request *FlightMostBookedDestinationsRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostBookedDestinations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("originCityCode", request.OriginCityCode)
	q.Add("period", string(request.Period))
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightBusiestTravelingPeriod(_ context.Context, request *FlightBusiestTravelingPeriodRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightBusiestTravelingPeriod)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("cityCode", request.CityCode)
	q.Add("period", string(request.Period))
	q.Add("direction", string(request.Direction))
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) AirportNearestRelevant(_ context.Context, request *AirportNearestRelevantRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightBusiestTravelingPeriod)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("latitude", fmt.Sprintf("%f", request.Latitude))
	q.Add("longitude", fmt.Sprintf("%f", request.Longitude))
	q.Add("sort", request.Sort)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) AirportAndCitySearch(_ context.Context, request *AirportAndCitySearchRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightBusiestTravelingPeriod)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("countryCode", request.CountryCode)
	q.Add("subType", request.SubType)
	q.Add("keyword", request.Keyword)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) AirlineCodeLookup(_ context.Context, request *AirlineCodeLookupRequest) (response *Response, err error) {
	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.AirlineCodeLookup)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("airlineCodes", request.AirlineCodes)
	req.URL.RawQuery = q.Encode()

	bearer := getBearer(aSrv.token)
	req.Header.Add("Authorization", bearer)
	req.Header.Add("Accept", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func NewBasicService(port int, configFilename string, urlsFilename string, logger log.Logger) (AmadeusService, error) {
	s, err := registerService("amadeus-go", port, time.Second*15)
	if err != nil {
		return nil, err
	}

	urls, err := getServicesURLs(urlsFilename)
	if err != nil {
		return nil, err
	}

	token, err := getTokenFromAmadeus(configFilename, urls)
	if err != nil {
		return nil, err
	}

	var srv AmadeusService
	aSrv := amadeusService{
		urls:         urls,
		token:        token,
		registerInfo: s,
	}

	srv = loggingMiddleware(logger)(aSrv)
	return srv, nil
}

// =============================================================================
type serviceMiddleware func(service AmadeusService) AmadeusService

type amadeusService struct {
	token        *amadeusToken
	urls         *serviceUrls
	registerInfo *serviceReg
}

type serviceUrls struct {
	ApiBaseUrl                      string
	AuthUrl                         string
	FlightLowFareSearch             string
	FlightInspirationSearch         string
	FlightCheapestDateSearch        string
	FlightMostSearchedDestinations  string
	FlightMostSearchedByDestination string
	FlightCheckInLists              string
	FlightMostTraveledDestinations  string
	FlightMostBookedDestinations    string
	FlightBusiestTravelingPeriod    string
	AirportNearestRelevant          string
	AirportAndCitySearch            string
	AirlineCodeLookup               string
}
