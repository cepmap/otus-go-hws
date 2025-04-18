package config

import (
	"github.com/cepmap/otus-go-hws/hw12_13_14_15_calendar/internal/logger"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Server struct {
	Host     string `mapstructure:"host" env:"SRV_HOST"`
	HTTPPort string `mapstructure:"httpPort" env:"SRV_HTTP_PORT"`
	GRPCPort string `mapstructure:"grpcPort" env:"SRV_GRPC_PORT"`
}

type Calendar struct {
	Log struct {
		Level string `mapstructure:"level" env:"LOG_LEVEL"`
	} `mapstructure:"log"`
	Storage      Storage `mapstructure:"storage"`
	Server       Server  `mapstructure:"server"`
	PrintVersion bool
}

var CalendarSettings *Calendar

func InitCalendarSettings() {
	defaultSettings := defaultCalendarSettings()
	CalendarSettings = &defaultSettings

	versionFlag := pflag.Bool("version", false, "version app")
	pflag.String("loglevel", "INFO", "log level app")
	pflag.String("config", "./configs/calendar-compose.yaml", "Path to configuration file")
	pflag.String("server_host", "0.0.0.0", "server hostname")
	pflag.String("server_httpport", "8080", "server http port")

	pflag.Parse()

	if *versionFlag {
		defaultSettings.PrintVersion = *versionFlag
		return
	}

	viper.BindPFlags(pflag.CommandLine)

	viper.SetConfigFile(viper.Get("config").(string))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("APP")

	if err := viper.ReadInConfig(); err != nil {
		logger.Error(err.Error())
	}

	if err := viper.Unmarshal(&CalendarSettings); err != nil {
		logger.Error(err.Error())
	}

	envLogLevel := viper.Get("LOG_LEVEL")
	if envLogLevel != nil {
		CalendarSettings.Log.Level = envLogLevel.(string)
	}
}

func defaultCalendarSettings() Calendar {
	return Calendar{
		Log: struct {
			Level string "mapstructure:\"level\" env:\"LOG_LEVEL\""
		}{Level: "DEBUG"},
		PrintVersion: false,
	}
}
