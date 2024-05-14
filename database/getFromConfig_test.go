package database

import (
	"lifeChecker/config"
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

var influxYml = `
type: Easy!
spec:
  token: InfluxAuthenticationToken
  url: http://localhost:8086
  org: myOrg
  bucket: myBucket
  measurement: services
`

func TestGetDatabaseFromConfig(t *testing.T) {
	t.Run("testing Influx", compInflux)
}

func compInflux(t *testing.T) {
	var config config.DatabaseConfig
	err := yaml.Unmarshal([]byte(influxYml), &config)
	assert.Nil(t, err)

	influx := newInfluxFromConfig(config)

	assert.Equal(t, influx.Influxdb_token, "InfluxAuthenticationToken")
	assert.Equal(t, influx.Influxdb_url, "http://localhost:8086")
	assert.Equal(t, influx.Influxdb_org, "myOrg")
	assert.Equal(t, influx.Influxdb_bucket, "myBucket")
	assert.Equal(t, influx.Influxdb_measurement, "services")
}
