package config

import (
	"github.com/bradrogan/banking/logger"
	"go.uber.org/zap"
)

var Db *dbConfig = &dbConfig{
	Database: database{
		Driver:       "mysql",
		User:         "root",
		Host:         "localhost",
		Port:         "3306",
		DatabaseName: "default",
	},
}

type database struct {
	Driver       string
	User         string
	Host         string
	Port         string
	DatabaseName string `mapstructure:"database_name"`
}

type dbConfig struct {
	Database database
}

func (db *dbConfig) ConfigInit() error {
	err := ConfigDriver.Unmarshal(db)

	if err != nil {
		logger.Fatal("Could not load db config", zap.Error(err))
		return err
	}
	return nil
}
