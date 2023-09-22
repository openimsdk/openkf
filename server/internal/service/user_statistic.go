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

package service

import (
	"context"
	"time"

	"github.com/openimsdk/openkf/server/internal/dal/dao"
	responseparams "github.com/openimsdk/openkf/server/internal/params/response"
)

// UserStatistic user statistics service.
type UserStatistic struct {
	Service

	UserStatisticDao *dao.UserStatisticDao
}

// NewUserStatisticService return new service with context.
func NewUserStatisticService(c context.Context) *UserStatistic {
	return &UserStatistic{
		Service: Service{
			ctx: c,
		},
		UserStatisticDao: dao.NewUserStatisticDao(),
	}
}

// ReadUserOnlineTimePerDay read user online time per day.
func (svc *UserStatistic) ReadUserOnlineTimePerDay(uid string, startTime, endTime int64) ([]*responseparams.UserStatisticItem, error) {
	// TODO: Need to specify the aggregation period and function to get accurate data.
	aggPeriod := "1d"
	aggFn := "sum"

	resultTimeItems := make([]*responseparams.UserStatisticItem, 0)

	// get online time sum
	onlineTimeItems, err := svc.UserStatisticDao.ReadMeasurementWithUserAction(
		dao.UserActionMeasurement,
		dao.UserOnlineAction,
		uid,
		startTime,
		endTime,
		aggPeriod,
		aggFn,
	)
	if err != nil {
		return resultTimeItems, err
	}

	offlineTimeItems, err := svc.UserStatisticDao.ReadMeasurementWithUserAction(
		dao.UserActionMeasurement,
		dao.UserOfflineAction,
		uid,
		startTime,
		endTime,
		aggPeriod,
		aggFn,
	)
	if err != nil {
		return resultTimeItems, err
	}

	// check data length
	// Maybe online times is longer than offline times.
	// if len(onlineTimeItems) != len(offlineTimeItems) {
	// 	return resultTimeItems, nil
	// }

	// merge online and offline time items (diff)
	for i := 0; i < len(offlineTimeItems); i++ {
		// Check is the same day
		onlineTime := time.Unix(onlineTimeItems[i].Time, 0)
		offlineTime := time.Unix(offlineTimeItems[i].Time, 0)
		isSameDay := onlineTime.Year() == offlineTime.Year() && onlineTime.Month() == offlineTime.Month() && onlineTime.Day() == offlineTime.Day()

		// skip if not the same day
		if !isSameDay {
			continue
		}

		// Skip today
		if offlineTimeItems[i].Time-onlineTimeItems[i].Time >= 24*60*60 {
			continue
		}

		resultTimeItems = append(resultTimeItems, &responseparams.UserStatisticItem{
			Data:      offlineTimeItems[i].Value - onlineTimeItems[i].Value,
			Timestamp: offlineTimeItems[i].Time,
		})
	}

	return resultTimeItems, nil
}

// ReadUserSendMsgCountPerDay read user send message count per day.
func (svc *UserStatistic) ReadUserSendMsgCountPerDay(uid string, startTime, endTime int64) ([]*responseparams.UserStatisticItem, error) {
	return svc.readUserSendMsgXXXPerDay(uid, startTime, endTime, dao.UserSendMsgCount)
}

// ReadUserSendMsgLenPerDay read user send message length per day.
func (svc *UserStatistic) ReadUserSendMsgLenPerDay(uid string, startTime, endTime int64) ([]*responseparams.UserStatisticItem, error) {
	return svc.readUserSendMsgXXXPerDay(uid, startTime, endTime, dao.UserSendMsgLen)
}

// readUserSendMsgXXXPerDay read user send message xxx metrics per day.
func (svc *UserStatistic) readUserSendMsgXXXPerDay(uid string, startTime, endTime int64, action dao.UserAction) ([]*responseparams.UserStatisticItem, error) {
	// TODO: Need to specify the aggregation period and function to get accurate data.
	aggPeriod := "1d"
	aggFn := "sum"

	resultTimeItems := make([]*responseparams.UserStatisticItem, 0)

	// get online time sum
	sendMsgCountTimeItems, err := svc.UserStatisticDao.ReadMeasurementWithUserAction(
		dao.UserMessageMeasurement,
		action,
		uid,
		startTime,
		endTime,
		aggPeriod,
		aggFn,
	)
	if err != nil {
		return resultTimeItems, err
	}

	// change data type
	for _, item := range sendMsgCountTimeItems {
		resultTimeItems = append(resultTimeItems, &responseparams.UserStatisticItem{
			Data:      item.Value,
			Timestamp: item.Time,
		})
	}

	return resultTimeItems, nil
}
