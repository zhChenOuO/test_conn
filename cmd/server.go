package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"test_conn/internal/configuration"
	pkgGrpc "test_conn/pkg/delivery/grpc"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/grpc"
	"gitlab.com/howmay/gopher/helper"
	"gitlab.com/howmay/gopher/zlog"

	"go.uber.org/fx"
)

// ServerCmd 是此程式的Service入口點
var ServerCmd = &cobra.Command{
	Run: run,
	Use: "server",
}

var Module = fx.Options(
	fx.Provide(grpc.StartServer),
)

type ServerConfig struct {
	fx.Out

	Log  *zlog.Config `mapstructure:"log"`
	GRPC *grpc.Config `mapstructure:"grpc"`
}

func run(command *cobra.Command, _ []string) {
	defer helper.Recover(command.Context())

	var cfg ServerConfig

	if err := configuration.Init(&cfg); err != nil {
		log.Fatal().Err(err).Msg("fail to init config")
	}

	zlog.New(cfg.Log)

	app := fx.New(
		fx.Supply(cfg),
		Module,
		pkgGrpc.Module,
		fx.Logger(&log.Logger),
	)

	ctx := log.Logger.WithContext(context.Background())
	exitCode := 0
	if err := app.Start(ctx); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(exitCode)
		return
	}

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	<-stopChan
	log.Info().Msg("main: shutting down server...")

	stopCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Error().Msg(err.Error())
	}

	os.Exit(exitCode)
}
