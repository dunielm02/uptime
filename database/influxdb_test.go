package database

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestInflux(t *testing.T) {
	godotenv.Load("../.env")
	t.Run("testing Connect function", testConnectToInfluxDB)

	t.Run("testing WriteTimeSerie functionality", testWriteTimeSerie)

}

func testConnectToInfluxDB(t *testing.T) {
	db := getTestDb()

	err := db.Connect()

	assert.Nil(t, err)

	err = db.CloseConnection()

	assert.Nil(t, err)
}

func testWriteTimeSerie(t *testing.T) {
	db := getTestDb()

	err := db.Connect()

	if err != nil {
		t.Error(err)
	}

	for i := range 10 {
		timeSerie := TimeSerie{
			Name:        "test2",
			RequestTime: time.Duration(i)*time.Millisecond + time.Millisecond,
			Alive:       false,
		}
		db.WriteTimeSerie(timeSerie)
	}

	doQuery(t, db, "request-time")
	doQuery(t, db, "alive")

}

func doQuery(t *testing.T, db Influx, field string) {
	queryApi := db.client.QueryAPI(db.Influxdb_org)

	res, err := queryApi.Query(context.Background(), fmt.Sprintf(`
		from(bucket: "%s")
			|> range(start: -10m)
			|> filter(fn: (r) => r._field=="%s")
	`, db.Influxdb_bucket, field))

	if err != nil {
		t.Error(err)
	}
	var cont = 0
	for res.Next() {
		cont++
		result := res.Record()
		fmt.Printf("%s = %v\n", result.Field(), result.Value())
		assert.Equal(t, field, result.Field())
	}
	fmt.Println(cont)
	if cont < 10 {
		t.Errorf("only %d were found", cont)
	}
}

func getTestDb() Influx {
	return Influx{
		client: nil,
		influxSpec: influxSpec{
			Influxdb_token:       os.Getenv("INFLUXDB_TOKEN"),
			Influxdb_url:         os.Getenv("INFLUXDB_URL"),
			Influxdb_org:         os.Getenv("INFLUXDB_ORG"),
			Influxdb_bucket:      os.Getenv("INFLUXDB_BUCKET"),
			Influxdb_measurement: os.Getenv("INFLUXDB_MEASUREMENT"),
		},
	}
}
