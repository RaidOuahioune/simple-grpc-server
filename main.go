package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"
)

func main() {

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8000))
	fmt.Println("Server started at localhost:8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	grpcServer.RegisterService(&TodoService_ServiceDesc, &TodoServiceServerImpelementation{})
	grpcServer.Serve(lis)

}
