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
	"encoding/json"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/gen"
)

// UserDispatchDao user dispatch dao.
type UserDispatchDao struct {
	Dao

	mu sync.Mutex
}

// NewUserDispatchDao return a user dispatch dao.
func NewUserDispatchDao() *UserDispatchDao {
	query := gen.Use(db.GetMysqlDB())
	cache := cache.Use(db.GetRedis())

	return &UserDispatchDao{
		Dao: Dao{
			ctx:   context.Background(),
			query: query,
			cache: &cache,
		},

		mu: sync.Mutex{},
	}
}

// EnqueueStringLPush enqueue string left push.
func (d *UserDispatchDao) EnqueueStringLPush(key, value string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	err := (*d.cache).LPush(d.ctx, key, value)
	if err != nil {
		return errors.Wrapf(err, "enqueue key: %s , value: %s, err: %s", key, value, err.Error())
	}

	return nil
}

// EnqueueStringRPush enqueue string right push.
func (d *UserDispatchDao) EnqueueStringRPush(key, value string) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	err := (*d.cache).RPush(d.ctx, key, value)
	if err != nil {
		return errors.Wrapf(err, "enqueue key: %s , value: %s, err: %s", key, value, err.Error())
	}

	return nil
}

// DequeueStringRPop dequeue string right pop.
func (d *UserDispatchDao) DequeueStringRPop(key string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	value, err := (*d.cache).RPop(d.ctx, key)
	if err != nil {
		return "", errors.Wrapf(err, "dequeue key: %s , err: %s", key, err.Error())
	}

	return value, nil
}

// DequeueStringLPop dequeue string right pop.
func (d *UserDispatchDao) DequeueStringLPop(key string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	value, err := (*d.cache).LPop(d.ctx, key)
	if err != nil {
		return "", errors.Wrapf(err, "dequeue key: %s , err: %s", key, err.Error())
	}

	return value, nil
}

// QueueLen get queue length.
func (d *UserDispatchDao) QueueLen(key string) (int64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	length, err := (*d.cache).LLen(d.ctx, key)
	if err != nil {
		return 0, errors.Wrapf(err, "enqueue key: %s , err: %s", key, err.Error())
	}

	return length, nil
}

// DequeueWithEnqueue dequeue with enqueue.
func (d *UserDispatchDao) DequeueWithEnqueue(key string) (string, error) {
	d.mu.Lock()
	defer d.mu.Unlock()

	value, err := (*d.cache).RPop(d.ctx, key)
	if err != nil {
		return "", errors.Wrapf(err, "dequeue with enqueue: %s , err: %s", key, err.Error())
	}

	err = (*d.cache).LPush(d.ctx, key, value)
	if err != nil {
		return "", errors.Wrapf(err, "dequeue with enqueue: %s , err: %s", key, err.Error())
	}

	return value, nil
}

// AddUser add user to zset.
func (d *UserDispatchDao) AddUser(key, value string) (int64, error) {
	res, err := (*d.cache).ZAdd(d.ctx, key, &redis.Z{
		Member: value,
		Score:  float64(time.Now().Unix()),
	})

	return res, err
}

// RemoveUser remove user from zset.
func (d *UserDispatchDao) RemoveUser(key, value string) (int64, error) {
	res, err := (*d.cache).ZRem(d.ctx, key, value)

	return res, err
}

// GetUser get user from zset top.
func (d *UserDispatchDao) GetUser(key string) (string, error) {
	// Get latest and set timestamp to now
	res, err := (*d.cache).ZRange(d.ctx, key, 0, 0)
	if err != nil || len(res) != 1 {
		return "", err
	}

	// Update timestamp
	value := res[0]
	_, err = d.AddUser(key, value)
	if err != nil {
		return "", err
	}

	return value, err
}

// SetSlackMap set staffId and slack channel id.
func (d *UserDispatchDao) SetSlackMap(key, customID, staffID, slackChannelID string) error {
	sMap := NewSlackMap(staffID, slackChannelID)

	// JSON encode
	value, err := json.Marshal(sMap)
	if err != nil {
		return errors.Wrapf(err, "Marshal user map: %s , err: %s", key, err.Error())
	}

	// SetToMap
	err = (*d.cache).HSet(d.ctx, key, customID, string(value))
	if err != nil {
		return errors.Wrapf(err, "SetToMap user map: %s , err: %s", key, err.Error())
	}

	return nil
}

// GetSlackMap get staffId and slack channel id.
func (d *UserDispatchDao) GetSlackMap(key, customID string) *SlackMap {
	// Get from map
	value, err := (*d.cache).HGet(d.ctx, key, customID)
	if err != nil {
		return &SlackMap{}
	}

	// JSON decode
	sMap := &SlackMap{}
	err = json.Unmarshal([]byte(value), sMap)
	if err != nil {
		return &SlackMap{}
	}

	return sMap
}

// GetSlackIDs get slack user ids.
func (d *UserDispatchDao) GetSlackIDs(key string) ([]string, error) {
	return (*d.cache).HKeys(d.ctx, key)
}
