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
	"errors"

	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/shomali11/slacker"
)

// USER_DISPATCH_QUEUE_KEY user dispatch queue key.
const USER_DISPATCH_QUEUE_KEY = "openkf:user_dispatch_queue"
const USER_SLACK_MAP_KEY = "openkf:user_slack_map"

// UserDispatchService user service.
type UserDispatchService struct {
	Service

	UserDispatchDao *dao.UserDispatchDao
	SysUserDao      *dao.SysUserDao
}

// NewUserDispatchService return new service with context.
func NewUserDispatchService(c context.Context) *UserDispatchService {
	return &UserDispatchService{
		Service: Service{
			ctx: c,
		},

		UserDispatchDao: dao.NewUserDispatchDao(),
		SysUserDao:      dao.NewSysUserDao(),
	}
}

// AddUser add user to enqueue.
func (s *UserDispatchService) AddUser(uuid string) error {
	_, err := s.UserDispatchDao.AddUser(USER_DISPATCH_QUEUE_KEY, uuid)

	return err
}

// GetUser get user and update timestamp.
func (s *UserDispatchService) GetUser() (string, error) {
	res, err := s.UserDispatchDao.GetUser(USER_DISPATCH_QUEUE_KEY)
	if err != nil {
		return "", errors.New("can not find a user")
	}

	// return a default value if res is empty
	if res == "" {
		u, err := s.SysUserDao.First()
		if err != nil {
			return "", err
		}

		return u.UUID, nil
	}

	return res, nil
}

// DeleteUser delete user from queue.
func (s *UserDispatchService) DeleteUser(uuid string) error {
	_, err := s.UserDispatchDao.RemoveUser(USER_DISPATCH_QUEUE_KEY, uuid)

	return err
}

// SetSlackMap set slack map.
func (s *UserDispatchService) SetSlackMap(customID, staffID string, botContext slacker.BotContext) error {
	return s.UserDispatchDao.SetSlackMap(USER_SLACK_MAP_KEY, customID, staffID, botContext.Event().ChannelID)
}

// GetSlackMap get slack map.
func (s *UserDispatchService) GetSlackMap(customID string) *dao.SlackMap {
	return s.UserDispatchDao.GetSlackMap(USER_SLACK_MAP_KEY, customID)
}

// GetStaffID get all staff id.
func (s *UserDispatchService) GetSlackIDs() ([]string, error) {
	return s.UserDispatchDao.GetSlackIDs(USER_SLACK_MAP_KEY)
}

// SlackUserFilter filter user by slack id.
func (s *UserDispatchService) SlackUserFilter(uid string) bool {
	ids, err := s.GetSlackIDs()
	if err != nil || len(ids) == 0 {
		return false
	}

	for _, id := range ids {
		if id == uid {
			return true
		}
	}

	return false
}
