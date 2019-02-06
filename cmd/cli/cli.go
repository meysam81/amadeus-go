package main

import (
	"amadeus-go/pkg/transports"

	"context"
	"flag"
	"google.golang.org/grpc"
	"log"
	"strings"
	"time"
)

func main() {
	fs := flag.NewFlagSet("amadeus-cli", flag.ExitOnError)
	var (
		grpcAddr = fs.String("grpc-addr", ":8000", "gRPC listener address")
	)

	ctx, _ := context.WithTimeout(context.TODO(), time.Second*1)
	conn, err := grpc.DialContext(ctx, *grpcAddr, grpc.WithInsecure())
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	srv := transports.NewGRPCClient(conn)
	name := "Meysam"
	resp, err := srv.Greeter(context.TODO(), name)
	if err != nil {
		panic(err)
	}

	log.Printf("%v ---> \n", name)
	log.Printf("%v <--- %v\n", strings.Repeat(" ", len(name)), resp)
}
