package grpc

import (
	"test_conn/proto"

	"google.golang.org/grpc"
)

// RegisterHandler ...
func RegisterHandler(s *grpc.Server, srv *Handler) {
	proto.RegisterTestConnServiceServer(s, srv)
}
