package utils

import (
	"amadeus-go/pkg/services"
	"amadeus-go/pkg/transports"

	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func SendReq(grpcAddr *string) {
	ctx, _ := context.WithTimeout(context.TODO(), time.Second*1)
	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	srv := transports.NewGRPCClient(conn)

	query := services.FlightLowFareSearchRequest{
		ReturnDate:    "2019-08-28",
		Destination:   "ELS",
		DepartureDate: "2019-08-27",
		Origin:        "NYC",
	}

	resp, err := srv.FlightLowFareSearch(context.TODO(), &query)
	if err != nil {
		panic(err)
	}

	log.Println(">>>>>", query)
	log.Println("<<<<<", resp)
}
