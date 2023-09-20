// Copyright © 2023 OpenIM open source community. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"fmt"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/pkg/log"
)

type InfluxDB struct {
	// for query
	Org    string
	Bucket string

	write api.WriteAPIBlocking
	query api.QueryAPI
}

var influxcient *InfluxDB

// InitInfluxDB init influxdb connection.
func InitInfluxDB() {
	severURL := fmt.Sprintf("http://%s:%d",
		config.Config.InfluxDB.Ip,
		config.Config.InfluxDB.Port,
	)

	influxDBClient := influxdb2.NewClient(severURL, config.Config.InfluxDB.Token)
	if influxDBClient == nil {
		log.Panic("InfluxDB", " open failed ", severURL)
	}

	_client := &InfluxDB{
		Org:    config.Config.InfluxDB.Org,
		Bucket: config.Config.InfluxDB.Bucket,
		write:  influxDBClient.WriteAPIBlocking(config.Config.InfluxDB.Org, config.Config.InfluxDB.Bucket),
		query:  influxDBClient.QueryAPI(config.Config.InfluxDB.Org),
	}

	influxcient = _client
	log.Info("Influx", "connect ok", severURL)
}

// GetInfluxDBClient get influxdb client.
func GetInfluxDBClient() *InfluxDB {
	return influxcient
}

// GetWriteAPI get write api.
func (i *InfluxDB) GetWriteAPI() api.WriteAPIBlocking {
	return i.write
}

// GetQueryAPI get query api.
func (i *InfluxDB) GetQueryAPI() api.QueryAPI {
	return i.query
}

// BuildQuery build query string.
func (i *InfluxDB) BuildQuery(bucket, measurement, field, aggPeriod, aggFn, actionStr, uuid string,
	startTime, endTime int64) string {
	return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: %d, stop: %d)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r._field == "%s")	
		|> filter(fn: (r) => r.uuid == "%s")
		|> filter(fn: (r) => r.action == "%s")
		|> aggregateWindow(every: %s, fn: %s, createEmpty: false)
		|> yield(name: "results")`, bucket, startTime, endTime, measurement,
		field, uuid, actionStr, aggPeriod, aggFn)
}

// BuildQueryWithPage build query string with page.
func (i *InfluxDB) BuildQueryWithPage(bucket, period, measurement, uuid string,
	pagesize, page int) string {
	// use group to get the all data
	return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r.uuid == "%s")
		|> group(columns: ["*"])
		|> sort(columns: ["_time"], desc: true)
		|> limit(n: %d, offset: %d)
		|> yield(name: "results")`, bucket, period, measurement,
		uuid, pagesize, (page-1)*pagesize)
}

// BuildQueryWithCount build query string with count.
func (i *InfluxDB) BuildQueryWithCount(bucket, period, measurement, uuid string) string {
	// use group to get the all data
	return fmt.Sprintf(`
		from(bucket: "%s")
		|> range(start: -%s)
		|> filter(fn: (r) => r._measurement == "%s")
		|> filter(fn: (r) => r.uuid == "%s")
		|> group(columns: ["*"])	
		|> count()`, bucket, period, measurement, uuid)
}
