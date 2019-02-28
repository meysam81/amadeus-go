package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

type authentication struct {
	ApiKey    string `json:"API_KEY"`
	ApiSecret string `json:"API_SECRET"`
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

func getTokenFromAmadeus(configFilename string, urls *serviceUrls) (*amadeusToken, error) {
	var (
		auth  authentication
		err   error
		token amadeusToken
	)

	err = readConf(configFilename, &auth)
	if err != nil {
		return nil, err
	}

	// this is the way to send body of mime-type: x-www-form-urlencoded
	body := url.Values{}
	body.Set("client_id", auth.ApiKey)
	body.Set("client_secret", auth.ApiSecret)
	body.Set("grant_type", "client_credentials")

	contentType := "application/x-www-form-urlencoded"
	urlStr := cleanUrl(urls.ApiBaseUrl, urls.AuthUrl)
	resp, err := http.Post(urlStr, contentType, strings.NewReader(body.Encode()))
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

func getServicesURLs(urlsFilename string) (*serviceUrls, error) {
	var urls serviceUrls

	err := readConf(urlsFilename, &urls)
	if err != nil {
		return nil, err
	}

	return &urls, nil
}

func cleanUrl(base, route string) string {
	return base + route
}

func getBearer(token *amadeusToken) string {
	return token.TokenType + " " + token.AccessToken
}

func readConf(filename string, config interface{}) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(&config)
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
