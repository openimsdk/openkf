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
	"time"

	"github.com/gin-gonic/gin"

	"github.com/openimsdk/openkf/server/internal/dal/dao"
	"github.com/openimsdk/openkf/server/pkg/log"
)

// IMCallbackService openim callback service.
type IMCallbackService struct {
	Service

	UserStatisticDao *dao.UserStatisticDao
}

// NewIMCallbackService return new service with gin context.
func NewIMCallbackService(c *gin.Context) *IMCallbackService {
	return &IMCallbackService{
		Service: Service{
			ctx: c,
		},
		UserStatisticDao: dao.NewUserStatisticDao(),
	}
}

// UserOnlineCallback user online callback.
func (svc *IMCallbackService) UserOnlineCallback(uid string) error {
	if err := svc.UserStatisticDao.WriteMeasurementWithUserActionData(
		uid,
		dao.UserActionMeasurement,
		dao.UserOnlineAction,
		time.Now().Unix(),
	); err != nil {
		log.Errorf("UserOnlineCallback", "Write to influx db error: %s", err.Error())

		return err
	}

	return nil
}

// UserOfflineCallback user offline callback.
func (svc *IMCallbackService) UserOfflineCallback(uid string) error {
	if err := svc.UserStatisticDao.WriteMeasurementWithUserActionData(
		uid,
		dao.UserActionMeasurement,
		dao.UserOfflineAction,
		time.Now().Unix(),
	); err != nil {
		log.Errorf("UserOnlineCallback", "Write to influx db error: %s", err.Error())

		return err
	}

	return nil
}

// AfterSendSingleMsgCallback after send msg callback.
func (svc *IMCallbackService) AfterSendSingleMsgCallback(uid string, msg string) error {
	if err := svc.UserStatisticDao.WriteMeasurementWithUserActionData(
		uid,
		dao.UserMessageMeasurement,
		dao.UserSendMsgLen,
		int64(len(msg)), // TODO: Check true length of the message
	); err != nil {
		log.Errorf("UserOnlineCallback", "Write to influx db error: %s", err.Error())

		return err
	}

	if err := svc.UserStatisticDao.WriteMeasurementWithUserActionData(
		uid,
		dao.UserMessageMeasurement,
		dao.UserSendMsgCount,
		1,
	); err != nil {
		log.Errorf("UserOnlineCallback", "Write to influx db error: %s", err.Error())

		return err
	}

	return nil
}
