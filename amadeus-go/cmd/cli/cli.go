package main

import (
	"Amadeus/amadeus-go/pkg/services"
	"Amadeus/amadeus-go/pkg/transports"

	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	fs := flag.NewFlagSet("amadeus-cli", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8080", "gRPC listener address")
	)

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
