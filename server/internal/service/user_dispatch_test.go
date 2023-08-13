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
	"testing"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// TestUserQueue test user queue.
func TestUserQueue(t *testing.T) {
	// Init
	config.ConfigInit("../../config.yaml")
	log.InitLogger()
	db.InitMysqlDB()
	db.InitRedisDB()

	s := NewUserDispatchService(context.Background())
	// err := s.AddUser("test3")
	// if err != nil {
	// 	t.Error(err)
	// }
	value, err := s.GetUser()
	if err != nil {
		t.Error(err)
	}
	t.Error(value)
}
