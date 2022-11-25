package config

import (
	"github.com/bradrogan/banking/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var ConfigDriver *viper.Viper

// Add config types here so they are loaded on config init
var config = [...]configInitter{
	App,
	Db,
}

type configInitter interface {
	ConfigInit() error
}

func InitConfiguration(filepath string) {
	ConfigDriver = viper.New()
	ConfigDriver.SetConfigFile(filepath)

	ConfigDriver.OnConfigChange(configChange)
	ConfigDriver.WatchConfig()

	err := ConfigDriver.ReadInConfig()
	if err != nil {
		logger.Panic("Could not read config", zap.Error(err))
	}

	if err := loadConfiguration(); err != nil {
		logger.Panic("could not load config", zap.Error(err))
	}
}

func configChange(e fsnotify.Event) {
	logger.Info("reloading config: ", zap.String("config_file", e.Name))
	if err := loadConfiguration(); err != nil {
		logger.Error("could not reload configuration", zap.String("load_errr", err.Error()))
	}
}

func loadConfiguration() error {

	for _, val := range config {
		if err := val.ConfigInit(); err != nil {
			return err
		}
	}
	return nil
}
