package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/kelseyhightower/envconfig"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

type authentication struct {
	ApiUrl    string `envconfig:"API_URL" required:"true"`
	ApiKey    string `envconfig:"API_KEY" required:"true"`
	ApiSecret string `envconfig:"API_SECRET" required:"true"`
}

type amadeusToken struct {
	Type        string        `json:"type"`
	Username    string        `json:"username"`
	AppName     string        `json:"application_name"`
	ClientId    string        `json:"client_id"`
	TokenType   string        `json:"token_type"`
	AccessToken string        `json:"access_token"`
	ExpiresIn   time.Duration `json:"expires_in"`
	State       string        `json:"state"`
	Scope       string        `json:"scope"`
}

func getTokenFromAmadeus() (*amadeusToken, error) {
	var (
		auth  authentication
		err   error
		token amadeusToken
	)

	err = envconfig.Process("auth", &auth)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: x-www-form-urlencoded
	body := url.Values{}
	body.Set("client_id", auth.ApiKey)
	body.Set("client_secret", auth.ApiSecret)
	body.Set("grant_type", "client_credentials")

	contentType := "application/x-www-form-urlencoded"
	resp, err := http.Post(auth.ApiUrl,
		contentType,
		strings.NewReader(body.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

func getServicesURLs() (*serviceUrls, error) {
	// TODO read from config file
	urls := serviceUrls{
		apiBaseUrl:                     "https://test.api.amadeus.com",
		flightLowFareSearch:            "/v1/shopping/flight-offers",
		flightInspirationSearch:        "/v1/shopping/flight-destinations",
		flightCheapestDateSearch:       "/v1/shopping/flight-dates",
		flightMostTraveledDestinations: "/v1/travel/analytics/air-traffic/traveled",
		flightMostBookedDestinations:   "/v1/travel/analytics/air-traffic/booked",
		flightBusiestTravelingPeriod:   "/v1/travel/analytics/air-traffic/busiest-period",
		airportNearestRelevant:         "/v1/reference-data/locations/airports",
		airportAndCitySearch:           "/v1/reference-data/locations",
	}
	return &urls, nil
}

func cleanUrl(base, route string) string {
	return base + route
}

func getBearer(token *amadeusToken) string {
	return token.TokenType + " " + token.AccessToken
}

/* TODO put this somewhere later (don't remove them before that)
err := initRedis()
if err != nil {
	panic(err)
}

func getToken(redisClient *redis.Client) (*amadeusToken, error) {
	token, err := readConfRedis(redisClient)
	if err != nil {
		log.Fatalln("<read config from redis>", err)

		token, err = getTokenFromAmadeus()
		if err != nil {
			return nil, err
		}

		err = writeToRedis(redisClient, token)
		if err != nil {
			log.Fatalln("<write to redis>", err)
		}
	}

	return token, nil
}

*/
