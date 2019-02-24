package main

import (
	"amadeus-go/cmd/cli/utils"
	"amadeus-go/pkg/services"

	"flag"
	"os"
)

var (
	FLIGHT_LOW_FARE_SEARCH = &services.FlightLowFareSearchRequest{
		ReturnDate:    "2019-08-28",
		Destination:   "ELS",
		DepartureDate: "2019-08-27",
		Origin:        "NYC",
	}
)

func main() {
	fs := flag.NewFlagSet("amadeus-cli", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8000", "gRPC listener address")
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	utils.SendReq(grpcAddr, FLIGHT_LOW_FARE_SEARCH)
}
