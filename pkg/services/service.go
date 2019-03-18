package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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

	if !emptyString(request.ReturnDate) {
		q.Add("returnDate", request.ReturnDate)
	}
	if !emptyString(request.ArrivalBy) {
		q.Add("arrivalBy", request.ArrivalBy)
	}
	if !emptyString(request.ReturnBy) {
		q.Add("returnBy", request.ReturnBy)
	}
	if request.Adults > 0 {
		num := strconv.Itoa(int(request.Adults))
		q.Add("adults", num)
	}
	if request.Children > 0 {
		num := strconv.Itoa(int(request.Children))
		q.Add("children", num)
	}
	if request.Infants > 0 {
		num := strconv.Itoa(int(request.Infants))
		q.Add("infants", num)
	}
	if request.Seniors > 0 {
		num := strconv.Itoa(int(request.Seniors))
		q.Add("seniors", num)
	}
	if request.TravelClass != nil {
		t := string(*request.TravelClass)
		q.Add("travelClass", t)
	}
	if !emptyString(request.IncludeAirlines) {
		q.Add("includeAirlines", request.IncludeAirlines)
	}
	if !emptyString(request.ExcludeAirlines) {
		q.Add("excludeAirlines", request.ExcludeAirlines)
	}
	if request.NonStop {
		q.Add("nonStop", "true")
	}
	if !emptyString(request.Currency) {
		q.Add("currency", request.Currency)
	}
	if request.MaxPrice > 0 {
		num := strconv.Itoa(int(request.MaxPrice))
		q.Add("maxPrice", num)
	}
	if request.Max > 0 && request.Max <= 250 {
		num := strconv.Itoa(int(request.Max))
		q.Add("max", num)
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightInspirationSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("origin", request.Origin)

	if !emptyString(request.DepartureDate) {
		q.Add("departureDate", request.DepartureDate)
	}
	if request.OneWay {
		q.Add("oneWay", "true")
	}
	if !emptyString(request.Duration) {
		q.Add("duration", request.Duration)
	}
	if request.NonStop {
		q.Add("nonStop", "true")
	}
	if request.MaxPrice > 0 {
		q.Add("maxPrice", string(request.MaxPrice))
	}
	if !emptyString(request.Currency) {
		q.Add("currency", request.Currency)
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightCheapestDateSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("origin", request.Origin)
	q.Add("destination", string(request.Destination))

	if !emptyString(request.DepartureDate) {
		q.Add("departureDate", request.DepartureDate)
	}
	if request.OneWay {
		q.Add("oneWay", "true")
	}
	if !emptyString(request.Duration) {
		q.Add("duration", request.Duration)
	}
	if request.NonStop {
		q.Add("nonStop", "true")
	}
	if request.MaxPrice > 0 {
		q.Add("maxPrice", string(request.MaxPrice))
	}
	if !emptyString(request.Currency) {
		q.Add("currency", request.Currency)
	}
	if request.ViewBy != nil {
		q.Add("viewBy", string(*request.ViewBy))
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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

	if request.Max > 0 {
		q.Add("max", string(request.Max))
	}
	if !emptyString(request.Fields) {
		q.Add("fields", request.Fields)
	}
	if request.PageLimit > 0 {
		q.Add("page[limit]", string(request.PageLimit))
	}
	if request.PageOffset > 0 {
		q.Add("page[offset]", string(request.PageOffset))
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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

	if !emptyString(request.Fields) {
		q.Add("fields", request.Fields)
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightCheckInLists)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()

	q.Add("airlineCode", request.AirlineCode)

	if !emptyString(request.Language) {
		q.Add("language", request.Language)
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostTraveledDestinations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()
	q.Add("originCityCode", request.OriginCityCode)
	q.Add("period", string(request.Period))

	if request.Max > 0 {
		q.Add("max", string(request.Max))
	}
	if !emptyString(request.Fields) {
		q.Add("fields", request.Fields)
	}
	if request.PageLimit > 0 {
		q.Add("page[limit]", string(request.PageLimit))
	}
	if request.PageOffset > 0 {
		q.Add("page[offset]", string(request.PageOffset))
	}
	if request.Sort != nil {
		q.Add("sort", string(*request.Sort))
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightMostBookedDestinations)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()

	q.Add("originCityCode", request.OriginCityCode)
	q.Add("period", string(request.Period))

	if request.Max > 0 {
		q.Add("max", string(request.Max))
	}
	if !emptyString(request.Fields) {
		q.Add("fields", request.Fields)
	}
	if request.PageLimit > 0 {
		q.Add("page[limit]", string(request.PageLimit))
	}
	if request.PageOffset > 0 {
		q.Add("page[offset]", string(request.PageOffset))
	}
	if request.Sort != nil {
		q.Add("sort", string(*request.Sort))
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

	url := cleanUrl(aSrv.urls.ApiBaseUrl, aSrv.urls.FlightBusiestTravelingPeriod)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: application/x-www-form-urlencoded
	q := req.URL.Query()

	q.Add("cityCode", request.CityCode)
	q.Add("period", string(request.Period))

	if request.Direction != nil {
		q.Add("direction", string(*request.Direction))
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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
	err = checkTokenExpiry(&aSrv)
	if err != nil {
		return nil, err
	}

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

func NewBasicService(port int, configFilename string, urlsFilename string,
	serviceName string, logger log.Logger) (AmadeusService, error) {
	s, err := registerService(serviceName, port, time.Second*15)
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
		urls:           urls,
		token:          token,
		registerInfo:   s,
		configFilename: configFilename,
		urlsFilename:   urlsFilename,
	}

	srv = loggingMiddleware(logger)(aSrv)
	return srv, nil
}

// =============================================================================
type serviceMiddleware func(service AmadeusService) AmadeusService

type amadeusService struct {
	token          *amadeusToken
	urls           *serviceUrls
	registerInfo   *serviceReg
	configFilename string
	urlsFilename   string
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

// =============================================================================
func emptyString(s string) bool {
	return strings.Compare(s, "") == 0
}
