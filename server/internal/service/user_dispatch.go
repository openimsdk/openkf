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

	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
)

// USER_DISPATCH_QUEUE_KEY user dispatch queue key.
const USER_DISPATCH_QUEUE_KEY = "openkf:user_dispatch_queue"

// UserDispatchService user service.
type UserDispatchService struct {
	Service

	UserDispatchDao *dao.UserDispatchDao
}

// NewUserDispatchService return new service with context.
func NewUserDispatchService(c context.Context) *UserDispatchService {
	return &UserDispatchService{
		Service: Service{
			ctx: c,
		},

		UserDispatchDao: dao.NewUserDispatchDao(),
	}
}

// AddUser add user to enqueue.
func (s *UserDispatchService) AddUser(uuid string) error {
	_, err := s.UserDispatchDao.AddUser(USER_DISPATCH_QUEUE_KEY, uuid)

	return err
}

// GetUser get user and update timestamp.
func (s *UserDispatchService) GetUser() (string, error) {
	return s.UserDispatchDao.GetUser(USER_DISPATCH_QUEUE_KEY)
}

// DeleteUser delete user from queue.
func (s *UserDispatchService) DeleteUser(uuid string) error {
	_, err := s.UserDispatchDao.RemoveUser(USER_DISPATCH_QUEUE_KEY, uuid)

	return err
}
