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
	"encoding/json"
	"fmt"
	"net"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/shomali11/slacker"
	"github.com/slack-go/slack"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
	customerroles "github.com/OpenIMSDK/OpenKF/server/internal/models/customer_roles"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/constant"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/msg"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// SlackService slack service.
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
			// UUID: fmt.Sprintf("%s%s", dao.SLACK_PERFIX, userId),
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

// GetSlackUser get slack user.
func (svc *SlackService) GetSlackUser(userId string) (*customerroles.CustomerSlack, error) {
	return svc.CustomerSlackDao.FindFirstByUUID(userId)
}

// SendMsg send message to openkf.
func (svc *SlackService) SendMsg(uid, question string, botContext slacker.BotContext) error {
	udSvc := NewUserDispatchService(svc.ctx)

	// Get default staff id
	sMap := udSvc.GetSlackMap(uid)
	staffId := sMap.StaffID

	if sMap.StaffID == "" || sMap.SlackChannelID == "" {
		// Get staff id
		tempStaffId, err := udSvc.GetUser()
		if err != nil {
			return err
		}

		// Store staff, writer to cache redis map
		err = udSvc.SetSlackMap(uid, tempStaffId, botContext)
		if err != nil {
			return errors.Wrapf(err, "set slack map failed")
		}
		staffId = tempStaffId
	}

	// Get OpenIM admin token
	uSvc := NewUserService(&gin.Context{}) // TODO: Change context to same
	token, err := uSvc.GetAdminToken()
	if err != nil {
		return errors.Wrapf(err, "get admin token failed")
	}

	// Get custom info
	customer, err := svc.CustomerSlackDao.FindFirstByUUID(uid)
	if err != nil {
		return errors.Wrapf(err, "find customer failed")
	}

	msgInfo := &request.MsgInfo{
		SendID:           uid,
		RecvID:           staffId,
		GroupID:          "",
		SenderNickname:   customer.Nickname,
		SenderFaceURL:    customer.Avatar,
		SenderPlatformID: constant.PLATFORMID_WEB,
		Content: &request.TextContent{
			Text: fmt.Sprintf("{\"content\":\"%s\"}", question),
		},
		ContentType:     constant.CONTENT_TYPE_TEXT,
		SessionType:     constant.SESSION_TYPE_SINGLE_CHAT,
		IsOnlineOnly:    false,
		NotOfflinePush:  false,
		OfflinePushInfo: &request.OfflinePushInfo{},
	}
	res, err := json.Marshal(msgInfo)
	if err != nil {
		return errors.Wrapf(err, "marshal msgInfo failed")
	}
	log.Debugf("msgInfo", string(res))
	host := fmt.Sprintf("http://%s", net.JoinHostPort(config.Config.OpenIM.Ip, fmt.Sprintf("%d", config.Config.OpenIM.ApiPort)))
	resp, err := msg.AdminSendMsg(msgInfo, "sendMsg:"+uid, host, token.Token)
	if err != nil {
		return errors.Wrapf(err, "send msg failed")
	}
	log.Debugf("AdminSendMsg", "Resp: %+v", resp)

	return nil
}
