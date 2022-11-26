package connections

import (
	"time"

	"github.com/bradrogan/banking/config"
	"github.com/jmoiron/sqlx"
)

func NewDbClient() *sqlx.DB {
	dataSource :=
		config.Connections.Database.User +
			"@tcp(" +
			config.Connections.Database.Host +
			":" +
			config.Connections.Database.Port +
			")/" +
			config.Connections.Database.DatabaseName

	client, err := sqlx.Open(config.Connections.Database.Driver, dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
