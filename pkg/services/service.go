package services

import (
	"context"
)

type AmadeusService interface {
	Greeter(context.Context, string) (string, error)
}

type amadeusService struct{}

func (amadeusService) Greeter(_ context.Context, name string) (message string, err error) {
	if len(name) == 0 {
		name = "Amadeus"
	}
	message = "Hello " + name
	return
}

func NewBasicService() AmadeusService {
	return new(amadeusService)
}
