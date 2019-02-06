package transports

import (
	pb "amadeus-go/pb/amadeus"
	"amadeus-go/pkg/endpoints"
	"amadeus-go/pkg/services"

	"context"
	"github.com/go-kit/kit/endpoint"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

func NewGRPCClient(conn *grpc.ClientConn) services.AmadeusService {
	var greeterEndpoint endpoint.Endpoint
	{
		greeterEndpoint = grpcTransport.NewClient(
			conn,
			"go.kit.srv.amadeus.AmadeusService",
			"Greeter",
			encodeGreetingRequest,
			decodeGreetingResponse,
			pb.Response{},
		).Endpoint()
	}

	return endpoints.EndpointSet{GreeterEndpoint: greeterEndpoint}
}

func encodeGreetingRequest(_ context.Context, request interface{}) (interface{}, error) {
	req := request.(endpoints.GreeterRequest)
	return &pb.Request{Name: req.Name}, nil
}

func decodeGreetingResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(*pb.Response)
	return endpoints.GreeterResponse{Message: resp.Message}, nil
}
