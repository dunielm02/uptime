package database

import (
	"context"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

type Influx struct {
	client          influxdb2.Client
	influxdb_token  string
	influxdb_url    string
	influxdb_org    string
	Influxdb_bucket string
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

func (db *Influx) WriteTimeSerie(name string, requestTime time.Time, alive bool) error {
	writeAPI := db.client.WriteAPIBlocking(db.influxdb_org, db.Influxdb_bucket)

	tags := map[string]string{
		"name": name,
	}

	fields := map[string]interface{}{
		"time":  requestTime,
		"alive": alive,
	}
	point := write.NewPoint("measurement1", tags, fields, time.Now())

	if err := writeAPI.WritePoint(context.Background(), point); err != nil {
		return err
	}

	return nil
}

func (db *Influx) Connect() {
	db.client = influxdb2.NewClient(db.influxdb_url, db.influxdb_org)
}

func (db *Influx) CloseConnection() {
	db.client.Close()
}
