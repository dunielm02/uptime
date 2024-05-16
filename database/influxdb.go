package database

import (
	"context"
	"lifeChecker/config"
	"log"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	"github.com/mitchellh/mapstructure"
)

type Influx struct {
	client influxdb2.Client
	influxSpec
}

type influxSpec struct {
	Influxdb_token       string `mapstructure:"token"`
	Influxdb_url         string `mapstructure:"url"`
	Influxdb_org         string `mapstructure:"org"`
	Influxdb_bucket      string `mapstructure:"bucket"`
	Influxdb_measurement string `mapstructure:"measurement"`
}

func newInfluxFromConfig(dbc config.DatabaseConfig) Influx {
	var spec influxSpec

	err := mapstructure.Decode(dbc.Spec, &spec)
	if err != nil {
		log.Fatal("error converting database spec: ", err)
	}

	return Influx{
		client:     nil,
		influxSpec: spec,
	}
}

func (db *Influx) WriteTimeSerieFromChannel(ctx context.Context, data <-chan TimeSerie) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case i := <-data:
			err := db.WriteTimeSerie(i)
			if err != nil {
				log.Println("there was an error writing to influxdb: ", err)
			}
		}
	}
}

func (db *Influx) WriteTimeSerie(serie TimeSerie) error {
	writeAPI := db.client.WriteAPIBlocking(db.Influxdb_org, db.Influxdb_bucket)

	tags := map[string]string{
		"name": serie.Name,
	}

	fields := map[string]interface{}{
		"time":  serie.RequestTime / time.Millisecond,
		"alive": serie.Alive,
	}

	point := write.NewPoint(db.Influxdb_measurement, tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		return err
	}

	return nil
}

func (db *Influx) Connect() error {
	db.client = influxdb2.NewClientWithOptions(
		db.Influxdb_url,
		db.Influxdb_token,
		influxdb2.DefaultOptions().SetBatchSize(20),
	)

	_, err := db.client.Health(context.Background())

	return err
}

func (db *Influx) CloseConnection() error {
	db.client.Close()
	return nil
}
