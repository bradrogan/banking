package main

import (
	"github.com/bradrogan/banking/app"
	"github.com/bradrogan/banking/logger"
)

func main() {
	logger.Info("Starting application...")
	app.Start()
}
