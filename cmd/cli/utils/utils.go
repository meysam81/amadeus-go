package utils

import (
	sv "amadeus-go/pkg/services"
	"amadeus-go/pkg/transports"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func SendReq(grpcAddr *string, request interface{}) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*1)
	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	srv := transports.NewGRPCClient(conn)

	var (
		response interface{}
	)

	if req, ok := request.(*sv.FlightLowFareSearchRequest); ok {
		response, err = srv.FlightLowFareSearch(context.TODO(), req)
	} else if req, ok := request.(*sv.FlightInspirationSearchRequest); ok {
		response, err = srv.FlightInspirationSearch(context.TODO(), req)
	}
	if err != nil {
		panic(err)
	}

	log.Println("sending\n\t\t", request)
	log.Println("receiving\n\t\t", response)
}
