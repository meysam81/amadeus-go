package transports

import (
	pb "amadeus-go/pb/amadeus"
	"amadeus-go/pkg/endpoints"

	"context"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	GreeterHandler grpcTransport.Handler
}

func (s *grpcServer) Greeter(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	_, resp, err := s.GreeterHandler.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.Response), nil
}

func NewGRPCServer(endpoints endpoints.EndpointSet) pb.AmadeusServiceServer {
	return &grpcServer{
		GreeterHandler: grpcTransport.NewServer(
			endpoints.GreeterEndpoint,
			decodeGreetingRequest,
			encodeGreetingResponse,
		),
	}
}

func decodeGreetingRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Request)
	return endpoints.GreeterRequest{Name: req.Name}, nil
}

func encodeGreetingResponse(_ context.Context, response interface{}) (interface{}, error) {
	resp := response.(endpoints.GreeterResponse)
	return &pb.Response{Message: resp.Message}, nil
}
