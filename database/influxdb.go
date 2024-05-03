package database

import (
	"context"
	"lifeChecker/config"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Influx struct {
	client influxdb2.Client
	influxSpec
}

type influxSpec struct {
	influxdb_token       string
	influxdb_url         string
	influxdb_org         string
	influxdb_bucket      string
	influxdb_measurement string
}

func newInfluxFromConfig(dbc config.DatabaseConfig) Influx {
	spec := dbc.Spec.(influxSpec)

	return Influx{nil, spec}
}

func (db *Influx) WriteTimeSerieFromChannel(ctx context.Context, data <-chan TimeSerie) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case i := <-data:
			db.WriteTimeSerie(i)
		}
	}
}

func (db *Influx) WriteTimeSerie(serie TimeSerie) error {
	writeAPI := db.client.WriteAPIBlocking(db.influxdb_org, db.influxdb_bucket)

	tags := map[string]string{
		"name": serie.Name,
	}

	fields := map[string]interface{}{
		"time":  serie.RequestTime / time.Millisecond,
		"alive": serie.Alive,
	}

	point := write.NewPoint(db.influxdb_measurement, tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		return err
	}

	return nil
}

func (db *Influx) Connect() error {
	db.client = influxdb2.NewClient(db.influxdb_url, db.influxdb_token)
	return nil
}

func (db *Influx) CloseConnection() error {
	db.client.Close()
	return nil
}
