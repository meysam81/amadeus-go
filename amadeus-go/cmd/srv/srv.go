package main

import (
	"amadeus-go/pkg/endpoints"
	"amadeus-go/pkg/services"
	"amadeus-go/pkg/transports"
	pb "api/amadeus/func"

	"flag"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"net"
	"os"
)

func main() {
	fs := flag.NewFlagSet("amadeus-srv", flag.ContinueOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8080", "gRPC listen address")
	)
	fs.Parse(os.Args[1:])

	srv, err := services.NewBasicService()
	if err != nil {
		panic(err)
	}
	var (
		endpointSet = endpoints.NewEndpointSet(srv)
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
