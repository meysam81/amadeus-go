package main

import (
	pb "amadeus-go/pb/amadeus"
	"amadeus-go/pkg/endpoints"
	"amadeus-go/pkg/services"
	"amadeus-go/pkg/transports"

	"flag"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("amadeus-srv", flag.ContinueOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8000", "gRPC listen address")
	)
	fs.Parse(os.Args[1:])

	var (
		srv         = services.NewBasicService()
		endpointSet = endpoints.New(srv)
		grpcServer  = transports.NewGRPCServer(endpointSet)
	)

	grpcListener, err := net.Listen("tcp", *grpcAddr)
	if err != nil {
		panic(err)
	}

	baseServer := grpc.NewServer()
	pb.RegisterAmadeusServiceServer(baseServer, grpcServer)
	log.Fatalln(baseServer.Serve(grpcListener))
}
