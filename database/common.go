package database

import (
	"context"
	"lifeChecker/config"
	"log"
	"time"
)

type DB interface {
	WriteTimeSerie(TimeSerie) error
	WriteTimeSerieFromChannel(context.Context, <-chan TimeSerie) error
	Connect() error
	CloseConnection() error
}

type TimeSerie struct {
	Name        string
	RequestTime time.Duration
	Alive       bool
}

func GetDatabaseFromConfig(cfg config.DatabaseConfig) DB {
	switch cfg.Type {
	case "influxdb":
		db := newInfluxFromConfig(cfg)
		return &db

	default:
		log.Fatal("Database Not Recognized")
	}
	return nil
}
