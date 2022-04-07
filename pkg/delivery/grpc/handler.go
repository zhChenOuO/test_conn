package grpc

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		New,
	),
	fx.Invoke(
		RegisterHandler,
	),
)

// Params ...
type Params struct {
	fx.In
}

// Handler gRPC Handler ...
type Handler struct {

}

// New gRPC 依賴注入
func New(p Params) (*Handler, error) {
	h := Handler{

	}

	return &h, nil
}
