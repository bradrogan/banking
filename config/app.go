package config

import (
	"github.com/bradrogan/banking/logger"
	"go.uber.org/zap"
)

var App *appConfig = &appConfig{
	Server: server{
		Host: "localhost",
		Port: "8000",
	},
}

type server struct {
	Host string
	Port string
}

type appConfig struct {
	Server server
}

func (app *appConfig) ConfigInit() error {
	err := ConfigDriver.Unmarshal(app)

	if err != nil {
		logger.Fatal("Could not load app config", zap.Error(err))
		return err
	}
	return nil
}
