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
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/pkg/log"
)

var realClient RealRedisClient

// RedisClient the interface defines the methods for interacting with Redis clients.
type RedisClient interface {
	// default
	Get(ctx context.Context, key string) *redis.StringCmd
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd
	TTL(ctx context.Context, key string) *redis.DurationCmd
	Expire(ctx context.Context, key string, ttl time.Duration) *redis.BoolCmd
	ExpireAt(ctx context.Context, key string, ttl time.Time) *redis.BoolCmd
	Exists(ctx context.Context, keys ...string) *redis.IntCmd
	Del(ctx context.Context, keys ...string) *redis.IntCmd
	Incr(ctx context.Context, key string) *redis.IntCmd
	Close() error

	// List
	LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	LPop(ctx context.Context, key string) *redis.StringCmd
	RPop(ctx context.Context, key string) *redis.StringCmd
	LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd
	LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd
	LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd
	LLen(ctx context.Context, key string) *redis.IntCmd
	LIndex(ctx context.Context, key string, index int64) *redis.StringCmd

	// tx
	Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	Pipeline() redis.Pipeliner
	TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error)
	TxPipeline() redis.Pipeliner

	// ZSet
	ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd
	ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd
	ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd

	// map
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HGet(ctx context.Context, key, field string) *redis.StringCmd
	HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd
	HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd
	HExists(ctx context.Context, key, field string) *redis.BoolCmd
	HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd
	HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	HLen(ctx context.Context, key string) *redis.IntCmd
	HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd
	HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd

	Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error
}

// RealRedisClient implemented the RedisClient interface and used a real Redis client.
type RealRedisClient struct {
	client *redis.Client
}

// Get get value.
func (r *RealRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	return r.client.Get(ctx, key)
}

// Set set key value pair to redis.
func (r *RealRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return r.client.Set(ctx, key, value, expiration)
}

// TTL get ttl of key.
func (r *RealRedisClient) TTL(ctx context.Context, key string) *redis.DurationCmd {
	return r.client.TTL(ctx, key)
}

// Expire expire key.
func (r *RealRedisClient) Expire(ctx context.Context, key string, ttl time.Duration) *redis.BoolCmd {
	return r.client.Expire(ctx, key, ttl)
}

// ExpireAt expire key at time.
func (r *RealRedisClient) ExpireAt(ctx context.Context, key string, ttl time.Time) *redis.BoolCmd {
	return r.client.ExpireAt(ctx, key, ttl)
}

// Exists check some keys are exis.
func (r *RealRedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Exists(ctx, keys...)
}

// Del delete key.
func (r *RealRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	return r.client.Del(ctx, keys...)
}

// Incr incr key.
func (r *RealRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	return r.client.Incr(ctx, key)
}

// Close close redis client.
func (r *RealRedisClient) Close() error {
	return r.client.Close()
}

// LPush left push value to list.
func (r *RealRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.LPush(ctx, key, values...)
}

// RPush right push value to list.
func (r *RealRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.RPush(ctx, key, values...)
}

// LPop left pop value from list.
func (r *RealRedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.LPop(ctx, key)
}

// RPop right pop value from list.
func (r *RealRedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	return r.client.RPop(ctx, key)
}

// LRange get list value from start to stop.
func (r *RealRedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.LRange(ctx, key, start, stop)
}

// LRem remove count value from list.
func (r *RealRedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {
	return r.client.LRem(ctx, key, count, value)
}

// LTrim trim list from start to stop.
func (r *RealRedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	return r.client.LTrim(ctx, key, start, stop)
}

// LLen get list length.
func (r *RealRedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.LLen(ctx, key)
}

// LIndex get index value.
func (r *RealRedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	return r.client.LIndex(ctx, key, index)
}

// Pipelined pipelined.
func (r *RealRedisClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.client.Pipelined(ctx, fn)
}

// Pipeline pipline.
func (r *RealRedisClient) Pipeline() redis.Pipeliner {
	return r.client.Pipeline()
}

// TxPipelined tx pipelined.
func (r *RealRedisClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {
	return r.client.TxPipelined(ctx, fn)
}

// TxPipeline tx pipeline.
func (r *RealRedisClient) TxPipeline() redis.Pipeliner {
	return r.client.TxPipeline()
}

// Watch watch.
func (r *RealRedisClient) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {
	return r.client.Watch(ctx, fn, keys...)
}

// ZAdd zset add.
func (r *RealRedisClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	return r.client.ZAdd(ctx, key, members...)
}

// ZRem zset remove.
func (r *RealRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	return r.client.ZRem(ctx, key, members...)
}

// ZRange zset range.
func (r *RealRedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	return r.client.ZRange(ctx, key, start, stop)
}

// HSet hset.
func (r *RealRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return r.client.HSet(ctx, key, values)
}

// HGet hget.
func (r *RealRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	return r.client.HGet(ctx, key, field)
}

// HGetAll hgetall.
func (r *RealRedisClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	return r.client.HGetAll(ctx, key)
}

// HDel hdel.
func (r *RealRedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	return r.client.HDel(ctx, key, fields...)
}

// HExists hexists.
func (r *RealRedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	return r.client.HExists(ctx, key, field)
}

// HIncrBy hincrby.
func (r *RealRedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	return r.client.HIncrBy(ctx, key, field, incr)
}

// HKeys hkeys.
func (r *RealRedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	return r.client.HKeys(ctx, key)
}

// HLen hlen.
func (r *RealRedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	return r.client.HLen(ctx, key)
}

// HMGet hmget.
func (r *RealRedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	return r.client.HMGet(ctx, key, fields...)
}

// HMSet hmset.
func (r *RealRedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	return r.client.HMSet(ctx, key, values...)
}

// InitRedisDB init redis client.
func InitRedisDB() {
	realClient = RealRedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Ip, config.Config.Redis.Port),
			Password: config.Config.Redis.Password, // no password set
			DB:       config.Config.Redis.Database, // use default DB
		}),
	}
	pong, err := realClient.client.Ping(context.Background()).Result()
	if err != nil {
		log.Panic("Redis", "connect ping failed, err:", err.Error())
	} else {
		log.Info("Redis", "connect ping response:", pong)
	}
}

// GetRedis get redis client instance.
func GetRedis() RedisClient {
	return &realClient
}

// CloseRedis close redis client instance.
func CloseRedis() {
	if realClient.client != nil {
		err := realClient.Close()
		if err != nil {
			log.Error("Redis", "close failed, err:", err.Error())
		}
	}
}
