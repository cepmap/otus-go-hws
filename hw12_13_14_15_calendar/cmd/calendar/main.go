package main

import (
	"context"
	flag "github.com/spf13/pflag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/app"
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage/memory"
	//sqlstorage "github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVarP(&configFile, "config", "c", "config.yaml", "config file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := NewConfig(configFile)

	logg := logger.New(config.Logger.Level)

	storage := memorystorage.New()

	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
