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
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	"github.com/gin-gonic/gin"
)

// SlackService bot service.
type SlackService struct {
	Service
}

// NewSlackService return new service with gin context.
func NewSlackService(c *gin.Context) *SlackService {
	return &SlackService{
		Service: Service{
			ctx: c,
		},
	}
}

// GetSlackConfig get slack config.
func (svc *SlackService) GetSlackConfig() (*responseparams.SlackConfigResponse, error) {
	return &responseparams.SlackConfigResponse{
		BotToken:          config.Config.Slack.BotToken,
		AppToken:          config.Config.Slack.AppToken,
		AppID:             config.Config.Slack.AppID,
		ClientID:          config.Config.Slack.ClientID,
		ClientSecret:      config.Config.Slack.ClientSecret,
		SigningSecret:     config.Config.Slack.SigningSecret,
		VerificationToken: config.Config.Slack.VerificationToken,
	}, nil
}
