package grpc

import (
	"context"
	"test_conn/proto"
)

func (h *Handler) Ping(ctx context.Context, req *proto.PingReq) (*proto.PingResp, error) {
	return &proto.PingResp{}, nil
}
