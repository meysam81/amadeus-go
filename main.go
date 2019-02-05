package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	//_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"log"
	"time"

	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type authentication struct {
	ApiUrl    string `envconfig:"API_URL" required:"true"`
	ApiKey    string `envconfig:"API_KEY" required:"true"`
	ApiSecret string `envconfig:"API_SECRET" required:"true"`
}

type tokenResponse struct {
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

var (
	// access token from amadeus
	token               tokenResponse
	redisClient         *redis.Client
	redisAccessTokenKey = "amadeus_access_token"
)

//var db *gorm.DB

func getToken() error {
	var auth authentication
	err := envconfig.Process("auth", &auth)
	if err != nil {
		panic(err)
	}

	body := url.Values{}
	body.Set("client_id", auth.ApiKey)
	body.Set("client_secret", auth.ApiSecret)
	body.Set("grant_type", "client_credentials")

	resp, err := http.Post(auth.ApiUrl, "application/x-www-form-urlencoded", strings.NewReader(body.Encode()))
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	r, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(r, &token)
	if err != nil {
		return err
	}

	return nil
}

func initRedis() error {
	redisClient = redis.NewClient(
		&redis.Options{
			Addr:     "redis-server:6379",
			Password: "",
			DB:       0,
		},
	)

	_, err := redisClient.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}

func readConfRedis() error {
	accessToken, err := redisClient.Get(redisAccessTokenKey).Result()
	if err != nil {
		return err
	}

	ttl, err := redisClient.TTL(redisAccessTokenKey).Result()
	if err != nil {
		return err
	}

	token.AccessToken = accessToken
	token.ExpiresIn = ttl
	return nil
}

func writeToRedis() error {
	_, err := redisClient.SetNX(redisAccessTokenKey, token.AccessToken, time.Second*token.ExpiresIn).Result()
	if err != nil {
		return err
	}

	return nil
}

func init() {
	err := initRedis()
	if err != nil {
		panic(err)
	}

	err = readConfRedis()
	if err != nil {
		log.Fatalln("key not found in redis:", err)

		err = getToken()
		if err != nil {
			panic(err)
		}

		err = writeToRedis()
		if err != nil {
			log.Fatalln("could not write to redis:", err)
		}
	}

	return
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, &token)
	})

	e.Logger.Fatal(e.Start(":8000"))
}
