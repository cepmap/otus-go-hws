package main

import (
	"context"
	"fmt"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/config"
	sqlstorage "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage/sql"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/app"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/server/http"
)

var configFile string
var AppConfig config.Config

//func init() {
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

	var storage app.Storage
	switch config.Settings.Storage.Type {
	case "memory":
		storage = &sqlstorage.Storage{}
	case "db":
		storage = &sqlstorage.Storage{}
	default:
		logger.Warn("unknown storage type: " + config.Settings.Storage.Type + ", using memory")
		storage = &sqlstorage.Storage{}
	}

	calendar := app.New(storage)

	server := internalhttp.NewServer(calendar, fmt.Sprintf("%v:%v", config.Settings.Server.Host, config.Settings.Server.Port))

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logger.Error("failed to stop http server: " + err.Error())
		}
	}()

	logger.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logger.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
