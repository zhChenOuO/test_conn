package client

import (
	"strings"

	toolGrpc "gitlab.com/howmay/gopher/grpc"
	"google.golang.org/grpc"
)

// Configs ...
type Configs struct {
	TestConn ClientConfig `mapstructure:"test_conn"`
}

// ClientConfig ...
type ClientConfig struct {
	Addr string `mapstructure:"addr"`
	Port string `mapstructure:"port"`
}

// AdvertiseAddr ...
func (config ClientConfig) AdvertiseAddr() string {
	return config.Addr + ":" + strings.Trim(config.Port, ":")
}

// GrpcClient ...
func (config ClientConfig) GrpcClient() (*grpc.ClientConn, error) {
	return toolGrpc.NewClient(config.AdvertiseAddr())
}
