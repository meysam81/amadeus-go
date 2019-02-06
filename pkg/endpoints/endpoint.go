package endpoints

import (
	"amadeus-go/pkg/services"
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

type EndpointSet struct {
	GreeterEndpoint endpoint.Endpoint
}

func (s EndpointSet) Greeter(ctx context.Context, name string) (string, error) {
	resp, err := s.GreeterEndpoint(ctx, GreeterRequest{name})
	if err != nil {
		return "", err
	}

	response := resp.(GreeterResponse)
	return response.Message, nil
}

func MakeGreeterEndpoint(srv services.AmadeusService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(GreeterRequest)
		if !ok {
			return nil, errors.New("bad request")
		}
		resp, err := srv.Greeter(ctx, req.Name)
		return GreeterResponse{resp}, err
	}
}

func New(srv services.AmadeusService) EndpointSet {
	var greeterEndpoint endpoint.Endpoint
	greeterEndpoint = MakeGreeterEndpoint(srv)
	return EndpointSet{GreeterEndpoint: greeterEndpoint}
}

type GreeterRequest struct {
	Name string
}

type GreeterResponse struct {
	Message string `json:"message"`
}
