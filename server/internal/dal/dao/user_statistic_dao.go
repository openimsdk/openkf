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

package dao

import (
	"context"
	"time"

	"github.com/openimsdk/openkf/server/internal/conn/db"
	"github.com/openimsdk/openkf/server/internal/dal/cache"
	"github.com/openimsdk/openkf/server/internal/dal/gen"
	"github.com/openimsdk/openkf/server/internal/utils"
	"github.com/openimsdk/openkf/server/pkg/log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

// UserMeasurement user measurement.
type UserMeasurement string

// measurements.
const (
	USER_ACTION_MEASUREMENT  UserMeasurement = "user_action"
	USER_MESSAGE_MEASUREMENT UserMeasurement = "user_message"
)

// UserAction user action.
type UserAction string

// tags.
const (
	USER_UUID           UserAction = "uuid"
	USER_ONLINE_ACTION  UserAction = "online"
	USER_OFFLINE_ACTION UserAction = "offline"
	USER_SEND_MSG_COUNT UserAction = "send_count"
	USER_RECV_MSG_COUNT UserAction = "recv_count"
	USER_SEND_MSG_LEN   UserAction = "send_len"
	USER_RECV_MSG_LEN   UserAction = "recv_len"
)

// UserStatisticDao user statistic dao.
type UserStatisticDao struct {
	Dao
}

// NewUserStatisticDao return a user statistic dao.
func NewUserStatisticDao() *UserStatisticDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())
	// TODO: package influxdb operation
	influxClient := db.GetInfluxDBClient()

	return &UserStatisticDao{
		Dao: Dao{
			ctx:          context.Background(),
			query:        query,
			cache:        &cache,
			influxClient: influxClient,
		},
	}
}

// UserData user data.
type UserData struct {
	Data int64 `json:"data"`
}

// UserTags user tags.
type UserTags struct {
	UUID   string `json:"uuid"`
	Action string `json:"action"`
}

// TimeItem time item.
type TimeItem struct {
	Value int64 `json:"value"`
	Time  int64 `json:"time"`
}

// WriteMeasurementWithUserActionData write measurement with action data.
func (d *UserStatisticDao) WriteMeasurementWithUserActionData(uid string, measurement UserMeasurement, action UserAction, data int64) error {
	now := time.Now()

	// Convert struct to map
	userFields, err := utils.StructToMapWithJSONTag(&UserData{Data: data})
	if err != nil {
		log.Error("WriteUserOnline", "Can not convert user data to map", err)

		return err
	}
	userTags, err := utils.StructToMapString(&UserTags{UUID: uid, Action: string(action)})
	if err != nil {
		log.Error("WriteUserOnline", "Can not convert user data to map", err)

		return err
	}
	// Flatten map with one depth
	userFieldsFlattened := utils.FlattenMap(userFields, ".")

	userPoint := influxdb2.NewPoint(
		string(measurement),
		userTags,
		userFieldsFlattened,
		now, // set timezone
	)

	return d.influxClient.GetWriteAPI().WritePoint(context.Background(), userPoint)
}

// ReadMeasurementWithUserAction read measurement with user action.
func (d *UserStatisticDao) ReadMeasurementWithUserAction(measurement UserMeasurement, action UserAction, uid string, startTime, endTime int64, aggPeriod string, aggFn string) ([]*TimeItem, error) {
	bucket := d.influxClient.Bucket
	measurementStr := string(measurement)
	actionStr := string(action)

	// TODO: Set to constant
	field := "data"
	// process result to map
	dataList := make([]*TimeItem, 0)

	query := d.influxClient.BuildQuery(bucket, measurementStr, field, aggPeriod, aggFn, actionStr, uid, startTime, endTime)
	result, err := d.influxClient.GetQueryAPI().Query(context.Background(), query)
	if err != nil {
		log.Errorf("influxClient.GetQueryAPI().Query", "Can not query data: %s", err.Error())

		return dataList, err
	}

	for result.Next() {
		// process data
		dataList = append(dataList, &TimeItem{
			Value: int64(result.Record().Value().(float64)),
			Time:  result.Record().Time().Unix(),
		})
	}

	return dataList, nil
}
