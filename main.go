package main

import (
	"flag"

	"github.com/bradrogan/banking/app"
	"github.com/bradrogan/banking/config"
	"github.com/bradrogan/banking/logger"
)

var FlagConfigPath = flag.String("cfg", "./config.toml", "Path to config file")

func main() {
	logger.Info("Starting application...")
	logger.Info("Loading config...")
	config.InitConfiguration(*FlagConfigPath)
	app.Start()
}
