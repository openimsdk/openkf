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
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"

	"github.com/openimsdk/openkf/server/internal/config"
	"github.com/openimsdk/openkf/server/pkg/log"
)

const (
	// RedisClientMockMode redis mock client mode
	RedisClientMockMode = 1
	// RedisClientRealMode redis real client mode
	RedisClientRealMode = 2
)

var (
	op         RedisClientOp
	realClient RealRedisClient
	mockClient MockRedisClient
)

// ClientMode redis mode type
type ClientMode uint8

// RedisClientOp option of GetRedis
type RedisClientOp struct {
	ClientMode ClientMode
}

// RedisClientOptions options func of GetRedis
type RedisClientOptions func(conf *RedisClientOp)

// WithMock option func for getting mock client
func WithMock() RedisClientOptions {

	return func(op *RedisClientOp) {
		op.ClientMode = RedisClientMockMode
	}
}

// RedisClient the interface defines the methods for interacting with Redis clients
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

// RealRedisClient implemented the RedisClient interface and used a real Redis client
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

// Exists check some keys are exis
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

// LIndex get index value
func (r *RealRedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {

	return r.client.LIndex(ctx, key, index)
}

// Pipelined pipelined.
func (r *RealRedisClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {

	return r.client.Pipelined(ctx, fn)
}

// Pipline pipline.
func (r *RealRedisClient) Pipeline() redis.Pipeliner {

	return r.client.Pipeline()
}

// TxPipelined tx pipelined
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

// MockRedisClient is a mock Redis client for test
type MockRedisClient struct {
	data map[string]string
}

// Get get the key
func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	value, ok := m.data[key]
	if !ok {

		return redis.NewStringResult("", redis.Nil)
	}

	return redis.NewStringResult(value, nil)
}

// Set set the key
func (m *MockRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	m.data[key] = fmt.Sprintf("%v", value)

	return redis.NewStatusResult("OK", nil)
}

// TTL get ttl of key.
func (m *MockRedisClient) TTL(ctx context.Context, key string) *redis.DurationCmd {

	return redis.NewDurationResult(0, redis.Nil)
}

// Expire expire key.
func (m *MockRedisClient) Expire(ctx context.Context, key string, ttl time.Duration) *redis.BoolCmd {

	return redis.NewBoolResult(false, nil)
}

// ExpireAt expire key at time.
func (m *MockRedisClient) ExpireAt(ctx context.Context, key string, ttl time.Time) *redis.BoolCmd {

	return redis.NewBoolResult(false, nil)
}

// Exists check some keys are exist.
func (m *MockRedisClient) Exists(ctx context.Context, keys ...string) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// Del delete key.
func (m *MockRedisClient) Del(ctx context.Context, keys ...string) *redis.IntCmd {
	for _, v := range keys {
		delete(m.data, v)
	}

	return redis.NewIntResult(1, nil)
}

// Incr incr key.
func (m *MockRedisClient) Incr(ctx context.Context, key string) *redis.IntCmd {
	if _, ok := m.data[key]; !ok {
		m.data[key] = "0"
	}
	val, _ := strconv.ParseInt(m.data[key], 10, 64)
	val++
	m.data[key] = strconv.Itoa(int(val))

	return redis.NewIntResult(val, nil)
}

// Close close redis client.
func (m *MockRedisClient) Close() error {

	return nil
}

// LPush left push value to list.
func (m *MockRedisClient) LPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if _, ok := m.data[key]; !ok {
		m.data[key] = ""
	}

	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = fmt.Sprintf("%v", v)
	}

	m.data[key] = strings.Join(strValues, " ") + " " + m.data[key]
	length := int64(len(strValues))

	return redis.NewIntResult(length, nil)
}

// RPush right push value to list.
func (m *MockRedisClient) RPush(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	if _, ok := m.data[key]; !ok {
		m.data[key] = ""
	}

	strValues := make([]string, len(values))
	for i, v := range values {
		strValues[i] = fmt.Sprintf("%v", v)
	}

	m.data[key] = m.data[key] + " " + strings.Join(strValues, " ")
	length := int64(len(strValues))

	return redis.NewIntResult(length, nil)
}

// LPop left pop value from list.
func (m *MockRedisClient) LPop(ctx context.Context, key string) *redis.StringCmd {
	value := ""
	if _, ok := m.data[key]; ok {
		values := strings.Split(m.data[key], " ")
		if len(values) > 0 {
			value = values[0]
			m.data[key] = strings.Join(values[1:], " ")
		}
	}

	return redis.NewStringResult(value, nil)
}

// RPop right pop value from list.
func (m *MockRedisClient) RPop(ctx context.Context, key string) *redis.StringCmd {
	value := ""
	if _, ok := m.data[key]; ok {
		values := strings.Split(m.data[key], " ")
		if len(values) > 0 {
			value = values[len(values)-1]
			m.data[key] = strings.Join(values[:len(values)-1], " ")
		}
	}

	return redis.NewStringResult(value, nil)
}

// LRange get list value from start to stop.
func (m *MockRedisClient) LRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	if _, ok := m.data[key]; !ok {

		return redis.NewStringSliceCmd(ctx, nil)
	}

	values := strings.Split(m.data[key], " ")
	length := int64(len(values))

	if start < 0 {
		start += length
	}
	if stop < 0 {
		stop += length
	}

	if start < 0 {
		start = 0
	} else if start >= length {
		start = length - 1
	}
	if stop < 0 {
		stop = 0
	} else if stop >= length {
		stop = length - 1
	}

	if start > stop {

		return redis.NewStringSliceCmd(ctx, nil)
	}

	return redis.NewStringSliceResult(values[start:stop+1], nil)
}

// LRem remove count value from list.
func (m *MockRedisClient) LRem(ctx context.Context, key string, count int64, value interface{}) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// LTrim trim list from start to stop.
func (m *MockRedisClient) LTrim(ctx context.Context, key string, start, stop int64) *redis.StatusCmd {
	if _, ok := m.data[key]; !ok {

		return redis.NewStatusResult("OK", nil)
	}

	values := strings.Split(m.data[key], " ")
	length := int64(len(values))

	if start < 0 {
		start += length
	}
	if stop < 0 {
		stop += length
	}

	if start < 0 {
		start = 0
	} else if start >= length {
		start = length - 1
	}
	if stop < 0 {
		stop = 0
	} else if stop >= length {
		stop = length - 1
	}

	if start > stop {
		m.data[key] = ""
	} else {
		m.data[key] = strings.Join(values[start:stop+1], " ")
	}

	return redis.NewStatusResult("OK", nil)
}

// LLen get list length.
func (m *MockRedisClient) LLen(ctx context.Context, key string) *redis.IntCmd {
	if _, ok := m.data[key]; !ok {

		return redis.NewIntResult(0, nil)
	}

	values := strings.Split(m.data[key], " ")

	return redis.NewIntResult(int64(len(values)), nil)
}

// LIndex get index value
func (m *MockRedisClient) LIndex(ctx context.Context, key string, index int64) *redis.StringCmd {
	if _, ok := m.data[key]; !ok {

		return redis.NewStringResult("", redis.Nil)
	}

	values := strings.Split(m.data[key], " ")
	length := int64(len(values))

	if index < 0 {
		index += length
	}

	if index < 0 || index >= length {

		return redis.NewStringResult("", redis.Nil)
	}

	return redis.NewStringResult(values[index], nil)
}

// Pipelined pipelined.
func (m *MockRedisClient) Pipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {

	return []redis.Cmder{}, nil
}

// Pipeline pipline.
func (m *MockRedisClient) Pipeline() redis.Pipeliner {

	return nil
}

// TxPipelined tx pipeline.
func (m *MockRedisClient) TxPipelined(ctx context.Context, fn func(redis.Pipeliner) error) ([]redis.Cmder, error) {

	return []redis.Cmder{}, nil
}

// TxPipeline tx pipeline.
func (m *MockRedisClient) TxPipeline() redis.Pipeliner {

	return nil
}

// Watch watch.
func (m *MockRedisClient) Watch(ctx context.Context, fn func(*redis.Tx) error, keys ...string) error {

	return nil
}

// ZAdd zset add.
func (m *MockRedisClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// ZRem zset remove.
func (m *MockRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// ZRange zset range.
func (m *MockRedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {

	return redis.NewStringSliceCmd(ctx, 0, 0)
}

// HSet hset.
func (m *MockRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {

	return redis.NewIntCmd(ctx, key, values)
}

// HGet hget.
func (m *MockRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {

	return redis.NewStringResult(string(redis.Nil), nil)
}

// HGetAll hgetall.
func (m *MockRedisClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {

	return redis.NewStringStringMapCmd(ctx, nil)
}

// HDel hdel.
func (m *MockRedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// HExists hexists.
func (m *MockRedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {

	return redis.NewBoolCmd(ctx, false)
}

// HIncrBy hincrby.
func (m *MockRedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// HKeys hkeys.
func (m *MockRedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {

	return redis.NewStringSliceCmd(ctx, nil)
}

// HLen hlen.
func (m *MockRedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {

	return redis.NewIntResult(0, nil)
}

// HMGet hmget.
func (m *MockRedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {

	return redis.NewSliceCmd(ctx, nil)
}

// HMSet hmset.
func (m *MockRedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {

	return redis.NewBoolCmd(ctx, nil)
}

// InitRedisDB init redis client.
func InitRedisDB() {
	op = RedisClientOp{ClientMode: RedisClientRealMode}
	realClient = RealRedisClient{
		client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", config.Config.Redis.Ip, config.Config.Redis.Port),
			Password: config.Config.Redis.Password, // no password set
			DB:       config.Config.Redis.Database, // use default DB
		}),
	}
	mockClient = MockRedisClient{
		data: make(map[string]string),
	}
	pong, err := realClient.client.Ping(context.Background()).Result()
	if err != nil {
		log.Panic("Redis", "connect ping failed, err:", err.Error())
	} else {
		log.Info("Redis", "connect ping response:", pong)
	}
}

// GetRedis get redis client instance by optional parameters.
func GetRedis(ops ...RedisClientOptions) RedisClient {
	for _, option := range ops {
		option(&op)
	}
	if op.ClientMode == RedisClientMockMode {
		op.ClientMode = RedisClientRealMode

		return &mockClient
	} else {

		return &realClient
	}
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
