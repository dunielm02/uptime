package database

import (
	"context"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type DB interface {
	WriteTimeSerie(TimeSerie) error
	WriteTimeSerieFromChannel(context.Context, <-chan TimeSerie) error
	Connect() error
	CloseConnection() error
}

type Influx struct {
	client          influxdb2.Client
	influxdb_token  string
	influxdb_url    string
	influxdb_org    string
	Influxdb_bucket string
}

type TimeSerie struct {
	Name        string
	RequestTime time.Duration
	Alive       bool
}

// NewInfluxFromYml Not Yet
func NewInfluxFromYml() Influx {
	return Influx{}
}

func NewInfluxFromEnv() Influx {
	return Influx{
		influxdb_token:  os.Getenv("INFLUXDB_TOKEN"),
		influxdb_url:    os.Getenv("INFLUXDB_URL"),
		influxdb_org:    os.Getenv("myOrg"),
		Influxdb_bucket: os.Getenv("INFLUXDB_BUCKET"),
	}
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
	writeAPI := db.client.WriteAPIBlocking(db.influxdb_org, db.Influxdb_bucket)

	tags := map[string]string{
		"name": serie.Name,
	}

	fields := map[string]interface{}{
		"time":  serie.RequestTime / time.Millisecond,
		"alive": serie.Alive,
	}

	point := write.NewPoint("measurement1", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		return err
	}

	return nil
}

func (db *Influx) Connect() error {
	db.client = influxdb2.NewClient(db.influxdb_url, db.influxdb_org)
	return nil
}

func (db *Influx) CloseConnection() error {
	db.client.Close()
	return nil
}
