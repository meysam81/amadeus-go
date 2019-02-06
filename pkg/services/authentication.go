package services

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/kelseyhightower/envconfig"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	//_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	redisAccessTokenKey = "amadeus_access_token"
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

func getToken() (token *amadeusToken, err error) {
	var auth authentication
	err = envconfig.Process("auth", &auth)
	if err != nil {
		return nil, err
	}

	body := url.Values{}
	body.Set("client_id", auth.ApiKey)
	body.Set("client_secret", auth.ApiSecret)
	body.Set("grant_type", "client_credentials")

	resp, err := http.Post(auth.ApiUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(r, token)
	if err != nil {
		return nil, err
	}

	return
}

func initRedis() (redisClient *redis.Client, err error) {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     "redis-server:6379",
			Password: "",
			DB:       0,
		},
	)

	_, err = redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}

	return
}

func readConfRedis(redisClient *redis.Client) (token *amadeusToken, err error) {
	accessToken, err := redisClient.Get(redisAccessTokenKey).Result()
	if err != nil {
		return nil, err
	}

	ttl, err := redisClient.TTL(redisAccessTokenKey).Result()
	if err != nil {
		return nil, err
	}

	token.AccessToken = accessToken
	token.ExpiresIn = ttl
	return
}

func writeToRedis(redisClient *redis.Client, token *amadeusToken) error {
	_, err := redisClient.SetNX(redisAccessTokenKey, token.AccessToken, time.Second*token.ExpiresIn).Result()
	if err != nil {
		return err
	}

	return nil
}

/* TODO put this somewhere later (don't remove them before that)
err := initRedis()
if err != nil {
	panic(err)
}

err = readConfRedis()
if err != nil {
	log.Fatalln("<read config from redis>", err)

	err = getToken()
	if err != nil {
		panic(err)
	}

	err = writeToRedis()
	if err != nil {
		log.Fatalln("could not write to redis:", err)
	}
}

*/
