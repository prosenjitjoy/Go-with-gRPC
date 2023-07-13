package main

import (
	"main/proto"
	"net"

	"google.golang.org/grpc"
)

type HelloServer struct {
	proto.GreetServiceServer
}

func main() {
	lis, _ := net.Listen("tcp", ":5000")
	grpcServer := grpc.NewServer()
	proto.RegisterGreetServiceServer(grpcServer, &HelloServer{})
	grpcServer.Serve(lis)
}
