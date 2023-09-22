package test

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

// MockRedisClient is a mock Redis client for test.
type MockRedisClient struct {
	data map[string]string
}

// Get get the key.
func (m *MockRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	value, ok := m.data[key]
	if !ok {
		return redis.NewStringResult("", redis.Nil)
	}

	return redis.NewStringResult(value, nil)
}

// Set set the key.
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

// LIndex get index value.
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

// ZAdd simulates the ZADD command.
func (m *MockRedisClient) ZAdd(ctx context.Context, key string, members ...*redis.Z) *redis.IntCmd {
	count := 0
	if _, exists := m.data[key]; !exists {
		m.data[key] = ""
	}
	for _, member := range members {
		m.data[key] = member.Member.(string)
		count++
	}

	return redis.NewIntResult(int64(count), nil)
}

// ZRem simulates the ZREM command.
func (m *MockRedisClient) ZRem(ctx context.Context, key string, members ...interface{}) *redis.IntCmd {
	count := 0
	if _, exists := m.data[key]; exists {
		for _, member := range members {
			if m.data[key] == member.(string) {
				delete(m.data, key)
				count++

				break
			}
		}
	}

	return redis.NewIntResult(int64(count), nil)
}

// ZRange simulates the ZRANGE command.
func (m *MockRedisClient) ZRange(ctx context.Context, key string, start, stop int64) *redis.StringSliceCmd {
	var members []string
	if _, exists := m.data[key]; exists {
		members = append(members, m.data[key])
	}

	return redis.NewStringSliceResult(members, nil)
}

// HSet simulates the HSET command.
func (m *MockRedisClient) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	count := 0
	if _, exists := m.data[key]; !exists {
		m.data[key] = ""
	}
	for i := 0; i < len(values); i += 2 {
		field := values[i].(string)
		value := values[i+1].(string)
		m.data[key+":"+field] = value
		count++
	}

	return redis.NewIntResult(int64(count), nil)
}

// HGet simulates the HGET command.
func (m *MockRedisClient) HGet(ctx context.Context, key, field string) *redis.StringCmd {
	if value, exists := m.data[key+":"+field]; exists {
		return redis.NewStringResult(value, nil)
	}

	return redis.NewStringResult("", redis.Nil)
}

// HGetAll simulates the HGETALL command.
func (m *MockRedisClient) HGetAll(ctx context.Context, key string) *redis.StringStringMapCmd {
	result := make(map[string]string)
	for k, v := range m.data {
		if strings.HasPrefix(k, key+":") {
			field := strings.TrimPrefix(k, key+":")
			result[field] = v
		}
	}

	return redis.NewStringStringMapResult(result, nil)
}

// HDel simulates the HDEL command.
func (m *MockRedisClient) HDel(ctx context.Context, key string, fields ...string) *redis.IntCmd {
	count := 0
	for _, field := range fields {
		if _, exists := m.data[key+":"+field]; exists {
			delete(m.data, key+":"+field)
			count++
		}
	}

	return redis.NewIntResult(int64(count), nil)
}

// HExists simulates the HEXISTS command.
func (m *MockRedisClient) HExists(ctx context.Context, key, field string) *redis.BoolCmd {
	_, exists := m.data[key+":"+field]

	return redis.NewBoolResult(exists, nil)
}

// HIncrBy simulates the HINCRBY command.
func (m *MockRedisClient) HIncrBy(ctx context.Context, key, field string, incr int64) *redis.IntCmd {
	if _, exists := m.data[key+":"+field]; exists {
		value, _ := strconv.ParseInt(m.data[key+":"+field], 10, 64)
		value += incr
		m.data[key+":"+field] = strconv.FormatInt(value, 10)

		return redis.NewIntResult(value, nil)
	}

	return redis.NewIntResult(0, nil)
}

// HKeys simulates the HKEYS command.
func (m *MockRedisClient) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	var fields []string
	for k := range m.data {
		if strings.HasPrefix(k, key+":") {
			fields = append(fields, strings.TrimPrefix(k, key+":"))
		}
	}

	return redis.NewStringSliceResult(fields, nil)
}

// HLen simulates the HLEN command.
func (m *MockRedisClient) HLen(ctx context.Context, key string) *redis.IntCmd {
	count := 0
	for k := range m.data {
		if strings.HasPrefix(k, key+":") {
			count++
		}
	}

	return redis.NewIntResult(int64(count), nil)
}

// HMGet simulates the HMGET command.
func (m *MockRedisClient) HMGet(ctx context.Context, key string, fields ...string) *redis.SliceCmd {
	var values []interface{}
	for _, field := range fields {
		if value, exists := m.data[key+":"+field]; exists {
			values = append(values, value)
		} else {
			values = append(values, nil)
		}
	}

	return redis.NewSliceResult(values, nil)
}

// HMSet simulates the HMSET command.
func (m *MockRedisClient) HMSet(ctx context.Context, key string, values ...interface{}) *redis.BoolCmd {
	for i := 0; i < len(values); i += 2 {
		field := values[i].(string)
		value := values[i+1].(string)
		m.data[key+":"+field] = value
	}

	return redis.NewBoolResult(true, nil)
}
