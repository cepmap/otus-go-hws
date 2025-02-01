package main

import (
	"context"
	"fmt"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/app"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/config"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage/sql"
)

var AppConfig config.Config

// func init() {
//	flag.StringVarP(&configFile, "config", "c", "config.yaml", "config file")
//}

func main() {
	if config.Settings.PrintVersion {
		printVersion()
		return
	}

	logger.SetLogLevel(config.Settings.Log.Level)
	logger.Debug(config.Settings.DebugMessage)

	AppConfig.Storage.Type = "memory"

	var usedStorage storage.Storage
	if config.Settings.Storage.Type == "sql" {
		usedStorage = &sqlstorage.Storage{}
	} else {
		usedStorage = &memorystorage.Storage{}
	}
	usedStorage.InitStorage()
	defer usedStorage.Close()

	calendar := app.New(usedStorage)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	httpSrv := internalhttp.NewServer(calendar,
		fmt.Sprintf("%v:%v", config.Settings.Server.Host, config.Settings.Server.HTTPPort))
	grpcSrv := internalgrpc.NewServer(calendar,
		fmt.Sprintf("%v:%v", config.Settings.Server.Host, config.Settings.Server.GRPCPort))

	go func() {
		if err := httpSrv.Start(ctx); err != nil {
			logger.Error("failed to start http server: " + err.Error())
			cancel()
			os.Exit(1)
		}
		logger.Info("http is running...")
	}()

	go func() {
		if err := grpcSrv.Start(); err != nil {
			logger.Error("failed to start grpc server: " + err.Error())
			cancel()
			os.Exit(1)
		}
		logger.Info("grpc is running...")
	}()

	<-ctx.Done()

	ctx, cancel = context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	grpcSrv.Stop()
	if err := httpSrv.Stop(ctx); err != nil {
		logger.Error("failed to stop http server: " + err.Error())
	}
}
