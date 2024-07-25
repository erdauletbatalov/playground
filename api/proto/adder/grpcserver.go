package adder

import (
	"context"

	"playground/api/proto"
)

// GRPCServer ...
type GRPCServer struct {
	proto.UnimplementedAdderServer
}

func (s *GRPCServer) Add(ctx context.Context, req *proto.AddRequest) (*proto.AddResponse, error) {
	return &proto.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
