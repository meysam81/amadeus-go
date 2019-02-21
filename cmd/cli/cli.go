package main

import (
	"amadeus-go/cmd/cli/utils"
	"os"

	"flag"
)

func main() {
	fs := flag.NewFlagSet("amadeus-cli", flag.ExitOnError)
	var (
		grpcAddr= fs.String("grpc-addr", ":8000", "gRPC listener address")
	)

	err := fs.Parse(os.Args[1:])
	if err != nil {
		panic(err)
	}

	utils.SendReq(grpcAddr)
}
