package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"playground/api/proto"
	"playground/api/proto/adder"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	proto.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
