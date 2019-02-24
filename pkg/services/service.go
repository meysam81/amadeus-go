package services

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kit/kit/log"
)

type AmadeusService interface {
	FlightLowFareSearch(context.Context, *FlightLowFareSearchRequest) (*FlightLowFareSearchResponse, error)
	FlightInspirationSearch(context.Context, *FlightInspirationSearchRequest) (*FlightInspirationSearchResponse, error)
}

func (aSrv amadeusService) FlightLowFareSearch(_ context.Context, request *FlightLowFareSearchRequest) (response *FlightLowFareSearchResponse, err error) {
	url := cleanUrl(aSrv.urls.apiBaseUrl, aSrv.urls.flightLowFareSearch)
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
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func (aSrv amadeusService) FlightInspirationSearch(_ context.Context, request *FlightInspirationSearchRequest) (response *FlightInspirationSearchResponse, err error) {
	url := cleanUrl(aSrv.urls.apiBaseUrl, aSrv.urls.flightInspirationSearch)
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
	err = json.Unmarshal(b, &response)
	if err != nil {
		return nil, err
	}

	return
}

func NewBasicService(port int, logger log.Logger) (AmadeusService, error) {
	s, err := RegisterService("amadeus-go", port, time.Second*15)
	if err != nil {
		return nil, err
	}

	token, err := getTokenFromAmadeus()
	if err != nil {
		return nil, err
	}

	urls, err := getServicesURLs()
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
	apiBaseUrl              string
	flightLowFareSearch     string
	flightInspirationSearch string
}
