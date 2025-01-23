package main

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
)

// Config При желании конфигурацию можно вынести в internal/config.
// Организация конфига в main принуждает нас сужать API компонентов, использовать
// при их конструировании только необходимые параметры, а также уменьшает вероятность циклической зависимости.

type Config struct {
	Logger struct {
		Level string `mapstructure:"level"`
	} `mapstructure:"logger"`
	Storage struct {
		Type string `mapstructure:"type"`
	} `mapstructure:"storage"`
	DB struct {
		Name     string `mapstructure:"name" env:"DB_NAME"`
		Host     string `mapstructure:"host" env:"DB_HOST"`
		User     string `mapstructure:"user" env:"DB_USER"`
		Password string `mapstructure:"password" env:"DB_PASSWORD"`
	} `mapstructure:"db"`
	Server struct {
		Host string `mapstructure:"host" env:"SRV_HOST"`
		Port string `mapstructure:"port" env:"SRV_PORT"`
	} `mapstructure:"server"`
}

func NewConfig(configFile string) Config {

	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()

	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	slog.Info("Using config file: " + viper.ConfigFileUsed())

	var config Config
	config.Logger.Level = viper.GetString("logger.level")

	config.Storage.Type = viper.GetString("storage.type")

	if config.Storage.Type == "db" {
		config.DB.Name = viper.GetString("db.name")
		config.DB.Host = viper.GetString("db.host")
		config.DB.User = viper.GetString("db.user")
		config.DB.Password = viper.GetString("db.password")
	}

	config.Server.Host = viper.GetString("server.host")

	config.Server.Port = viper.GetString("server.port")

	fmt.Println(config.Storage.Type)

	return config
}

// TODO
