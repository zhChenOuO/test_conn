package client

import (
	"test_conn/proto"

	"go.uber.org/fx"
)

// Module ...
var Module = fx.Provide(
	NewTestConn,
	New,
)

// Params ...
type Params struct {
	fx.In

	TestConn proto.TestConnServiceClient
}

// New ...
func New(params Params) *GrpcClient {
	client := GrpcClient{
		TestConn: params.TestConn,
	}

	return &client
}

// GrpcClient ...
type GrpcClient struct {
	TestConn proto.TestConnServiceClient `yaml:"test_conn" mapstructure:"test_conn"`
}

// NewTestConn ...
func NewTestConn(config *Configs) (proto.TestConnServiceClient, error) {
	client, err := config.TestConn.GrpcClient()
	if err != nil {
		return nil, err
	}
	return proto.NewTestConnServiceClient(client), nil
}
