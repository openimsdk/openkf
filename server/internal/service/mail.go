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
	"time"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/cache"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
)

// MailService mail service.
type MailService struct {
	Service

	cache cache.Cache
}

// NewMailService return new service with gin context.
func NewMailService(c *gin.Context) *MailService {
	return &MailService{
		Service: Service{
			ctx: c,
		},
		cache: cache.Use(db.GetRedis()),
	}
}

// SendCode send code to email.
func (svc *MailService) SendCode(email string) (err error) {
	// Check the code is exist.
	code, err := svc.cache.Get(svc.ctx, "code:"+email)
	// Refresh code.
	if err != nil || code == "" {
		code = utils.GenerateCode()
		// save code in 60s
		err = svc.cache.Set(svc.ctx, "code:"+email, code, time.Second*60)

		return err
	}

	// Generate code.
	// err = client.SendEmail(email, "OpenKF Admin Register", "Your verification code is "+code)

	return err
}

// CheckCode check code.
func (svc *MailService) CheckCode(email, code string) bool {
	// Check the code is exist.
	c, err := svc.cache.Get(svc.ctx, "code:"+email)
	if err != nil {
		return false
	}

	if c == code {
		return true
	}

	return false
}
