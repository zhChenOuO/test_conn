package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"test_conn/internal/client"
	"test_conn/internal/configuration"
	"test_conn/proto"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"gitlab.com/howmay/gopher/helper"
	"gitlab.com/howmay/gopher/zlog"

	"go.uber.org/fx"
)

// ClientCmd 是此程式的Service入口點
var ClientCmd = &cobra.Command{
	Run: runClient,
	Use: "client",
}

type ClientConfig struct {
	fx.Out

	Log    *zlog.Config    `mapstructure:"log"`
	Client *client.Configs `mapstructure:"client"`
}

func runClient(command *cobra.Command, _ []string) {
	defer helper.Recover(command.Context())

	var cfg ClientConfig

	if err := configuration.Init(&cfg); err != nil {
		log.Fatal().Err(err).Msg("fail to init config")
	}

	zlog.New(cfg.Log)
	var testClient proto.TestConnServiceClient
	app := fx.New(
		fx.Supply(cfg),
		client.Module,
		fx.Logger(&log.Logger),
		fx.Populate(&testClient),
	)

	ctx := log.Logger.WithContext(context.Background())
	exitCode := 0
	if err := app.Start(ctx); err != nil {
		log.Error().Msg(err.Error())
		os.Exit(exitCode)
		return
	}

	go func() {
		for {
			if _, err := testClient.Ping(ctx, &proto.PingReq{}); err != nil {
				log.Err(err).Msg("fail to ping")
			}
			log.Info().Msg("ping done")
			time.Sleep(5 * time.Second)
		}
	}()

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
