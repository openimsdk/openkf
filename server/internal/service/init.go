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

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

// Service service
type Service struct {
	// Context
	Ctx context.Context

	// DB
	Dao   *gorm.DB
	Cache *redis.Client
}

// NewService return new service with context
func NewService() *Service {
	return &Service{
		Ctx:   context.Background(),
		Dao:   db.GetMysqlDB(),
		Cache: db.GetRedis(),
	}
}

// NewServiceWithGin return new service with gin context
func NewServiceWithGin(c *gin.Context) *Service {
	return &Service{
		Ctx:   c.Request.Context(),
		Dao:   db.GetMysqlDB(),
		Cache: db.GetRedis(),
	}
}
