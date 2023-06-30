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
	"context"
	"fmt"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/go-redis/redis/v8"
)

var r *redis.Client

// InitRedisDB init redis client
func InitRedisDB() {
	r = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Ip, config.Config.Redis.Port),
		Password: config.Config.Redis.Password, // no password set
		DB:       config.Config.Redis.Database, // use default DB
	})
	pong, err := r.Ping(context.Background()).Result()
	if err != nil {
		log.Panic("Redis", "connect ping failed, err:", err.Error())
	} else {
		log.Info("Redis", "connect ping response:", pong)
	}
}

// GetRedis get redis client instance
func GetRedis() *redis.Client {
	return r
}

// CloseRedis close redis client instance
func CloseRedis() {
	if r != nil {
		err := r.Close()
		if err != nil {
			log.Error("Redis", "close failed, err:", err.Error())
		}
	}
}
