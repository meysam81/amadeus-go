package main

import (
	pb "amadeus-go/api/amadeus/func"
	"amadeus-go/pkg/endpoints"
	"amadeus-go/pkg/services"
	"amadeus-go/pkg/transports"

	"flag"
	defaultLogger "log"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {
	fs := flag.NewFlagSet("amadeus-srv", flag.ContinueOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8000", "gRPC listen address")
	)
	fs.Parse(os.Args[1:])

	var port int
	var err error
	if strings.HasPrefix(*grpcAddr, ":") {
		port, err = strconv.Atoi((*grpcAddr)[1:])
	} else {
		port, err = strconv.Atoi((*grpcAddr)[1:])
	}
	if err != nil {
		panic(err)
	}

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "caller", log.DefaultCaller)

	srv, err := services.NewBasicService(port, logger)
	if err != nil {
		panic(err)
	}
	var (
		endpointSet = endpoints.NewEndpointSet(srv, logger)
		grpcServer  = transports.NewGRPCServer(endpointSet, logger)
	)

	grpcListener, err := net.Listen("tcp", string(*grpcAddr))
	if err != nil {
		panic(err)
	}

	baseServer := grpc.NewServer()
	pb.RegisterAmadeusServiceServer(baseServer, grpcServer)
	defaultLogger.Fatalln(
		"listening to", *grpcAddr,
		"serving...", baseServer.Serve(grpcListener),
	)
}
