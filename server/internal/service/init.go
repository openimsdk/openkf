// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package service

import (
	"context"

	"github.com/OpenIMSDK/OpenKF/server/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Service struct {
	// Context
	Ctx context.Context

	// DB
	Dao   *gorm.DB
	Cache *redis.Client
}

func NewService() *Service {
	return &Service{
		Ctx:   context.Background(),
		Dao:   db.GetMysqlDB(),
		Cache: db.GetRedis(),
	}
}

func NewServiceWithGin(c *gin.Context) *Service {
	return &Service{
		Ctx:   c.Request.Context(),
		Dao:   db.GetMysqlDB(),
		Cache: db.GetRedis(),
	}
}
