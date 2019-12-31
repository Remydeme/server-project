package db

import (
	"github.com/Remydeme/esme-devops-project/config"
	"github.com/go-bongo/bongo"
	"log"
)

var (
	Session *bongo.Connection
)

func init() {
	Session = NewDatabase(&config.Main.Database)
}

func NewDatabase(config *config.Database) *bongo.Connection {
	configMongo := &bongo.Config{
		ConnectionString: config.Host,
		Database:         config.DBname,
	}

	connection, err := bongo.Connect(configMongo)

	if err != nil {
		log.Fatal(err)
		panic("Enable to connect to database")
	}

	return connection
}
