package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	token := os.Getenv("INFLUXDB_TOKEN")
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)

	org := "myOrg"
	bucket := "firstBucket"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	for value := 0; value < 200; value++ {
		tags := map[string]string{
			"tagname1": "tagvalue1",
		}
		fields := map[string]interface{}{
			"field1": rand.Int() % 100,
			"field2": rand.Int() % 100,
		}
		point := write.NewPoint("measurement1", tags, fields, time.Now())
		time.Sleep(1 * time.Second) // separate points by 1 second

		if err := writeAPI.WritePoint(context.Background(), point); err != nil {
			log.Fatal(err)
		}
	}

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "firstBucket")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record().ValueByKey("_value"), results.Record().Field())
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}

	query = `from(bucket: "firstBucket")
              |> range(start: -100m)
              |> filter(fn: (r) => r._measurement == "measurement1")
              |> mean()`
	results, err = queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		fmt.Println(results.Record().ValueByKey("_value"))
	}
	if err := results.Err(); err != nil {
		log.Fatal(err)
	}
}
