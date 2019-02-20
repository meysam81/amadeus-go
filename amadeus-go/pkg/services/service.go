package services

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"io/ioutil"
	"net/http"
	"os"
)

type AmadeusService interface {
	FlightLowFareSearch(context.Context, *FlightLowFareSearchRequest) (*FlightLowFareSearchResponse, error)
}

type amadeusService struct {
	token *amadeusToken
	urls  *serviceUrls
}

type serviceUrls struct {
	apiBaseUrl          string
	flightLowFareSearch string
}

func (aSrv amadeusService) FlightLowFareSearch(_ context.Context, routeData *FlightLowFareSearchRequest) (*FlightLowFareSearchResponse, error) {
	url := cleanUrl(aSrv.urls.apiBaseUrl, aSrv.urls.flightLowFareSearch)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: form-data
	q := req.URL.Query()
	q.Add("origin", routeData.Origin)
	q.Add("destination", routeData.Destination)
	q.Add("departureDate", routeData.DepartureDate)
	q.Add("returnDate", routeData.ReturnDate)
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

	var flightOfferResult FlightLowFareSearchResponse
	b, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(b, &flightOfferResult)
	if err != nil {
		return nil, err
	}

	return &flightOfferResult, nil
}

func NewBasicService() (AmadeusService, error) {
	token, err := getTokenFromAmadeus()
	if err != nil {
		return nil, err
	}

	urls, err := getServicesURLs()
	if err != nil {
		return nil, err
	}

	var srv AmadeusService

	var aSrv amadeusService
	aSrv.token = token
	aSrv.urls = urls

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller)

	srv = loggingMiddleware(logger)(aSrv)

	return srv, nil
}

type ServiceMiddleware func(service AmadeusService) AmadeusService