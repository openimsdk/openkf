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
	"fmt"

	"github.com/slack-go/slack"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
	customerroles "github.com/OpenIMSDK/OpenKF/server/internal/models/customer_roles"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// SLACK_PERFIX do not use separator.
const SLACK_PERFIX = "slack"

// SlackService bot service.
type SlackService struct {
	Service

	CustomerSlackDao *dao.CustomerSlackDao
}

// NewSlackService return new service with Background context.
func NewSlackService(ctx context.Context) *SlackService {
	return &SlackService{
		Service: Service{
			ctx: ctx,
		},
		CustomerSlackDao: dao.NewCustomerSlackDao(),
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

// CreateCustomer create customer if not exists.
func (svc *SlackService) CreateCustomer(userId string, profile *slack.UserProfile) (string, uint, error) {
	// check if exists
	temp, err := svc.CustomerSlackDao.FindFirstByUUID(userId)
	if temp != nil || err == nil {
		return temp.UUID, temp.Id, err
	}

	// Set unique email
	if profile.Email == "" {
		profile.Email = fmt.Sprintf("%s %s", "NOEMAIL", utils.GenUUIDWithoutHyphen())
	}

	customerSlack := customerroles.CustomerSlack{
		UserBase: base.UserBase{
			// UUID:        fmt.Sprintf("%s%s", SLACK_PERFIX, userId),
			UUID:        userId,
			Email:       profile.Email,
			Nickname:    profile.FirstName,
			Avatar:      profile.Image512,
			Description: profile.Title,
			IsEnable:    true,
		},
		FirstName:             profile.FirstName,
		LastName:              profile.LastName,
		RealName:              profile.RealName,
		RealNameNormalized:    profile.RealNameNormalized,
		DisplayName:           profile.DisplayName,
		DisplayNameNormalized: profile.DisplayNameNormalized,
		Skype:                 profile.Skype,
		Phone:                 profile.Phone,
		Image24:               profile.Image24,
		Image32:               profile.Image32,
		Image48:               profile.Image48,
		Image72:               profile.Image72,
		Image192:              profile.Image192,
		Image512:              profile.Image512,
		ImageOriginal:         profile.ImageOriginal,
		Title:                 profile.Title,
		BotID:                 profile.BotID,
		ApiAppID:              profile.ApiAppID,
		StatusText:            profile.StatusText,
		StatusEmoji:           profile.StatusEmoji,
		StatusExpiration:      profile.StatusExpiration,
		Team:                  profile.Team,
	}
	if err = svc.CustomerSlackDao.Create(&customerSlack); err != nil {
		return customerSlack.UUID, 0, err
	}

	s, _ := svc.CustomerSlackDao.FindFirstByUUID(userId)

	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   userId,
				Nickname: customerSlack.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.CustomerSlackDao.Delete(s)

		return userId, s.Id, err
	}

	return userId, s.Id, nil
}
