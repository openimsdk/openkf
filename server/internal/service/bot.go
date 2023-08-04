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
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// BotService bot service.
type BotService struct {
	Service

	SysBotDao       *dao.SysBotDao
	SysCommunityDao *dao.SysCommunityDao
}

// NewBotService return new service with gin context.
func NewBotService(c *gin.Context) *BotService {
	return &BotService{
		Service: Service{
			ctx: c,
		},
		SysBotDao:       dao.NewSysBotDao(),
		SysCommunityDao: dao.NewSysCommunityDao(),
	}
}

// CreateBot create bot.
func (svc *BotService) CreateBot(cid string, params *requestparams.CreateBotParams) (string, uint, error) {
	// Create bot
	uid := utils.GenUUID()

	// Get community id
	_uuid, err := uuid.FromString(cid)
	if err != nil {
		return "", 0, err
	}
	communityInfo, err := svc.SysCommunityDao.FindFirstByUUID(_uuid)
	if err != nil {
		return "", 0, err
	}

	bot := &systemroles.SysBot{
		UUID:        uid,
		BotAddr:     params.BotAddr,
		BotPort:     params.BotPort,
		BotToken:    params.BotToken,
		Nickname:    params.Nickname,
		Avatar:      *params.Avatar,
		Description: *params.Description,
		CommunityId: communityInfo.Id,
	}
	if err = svc.SysBotDao.Create(bot); err != nil {
		return uid.String(), 0, err
	}

	b, _ := svc.SysBotDao.FindFirstByUUID(uid)

	// Register bot to OpenIM
	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   uid.String(),
				Nickname: params.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.SysBotDao.Delete(b)

		return uid.String(), b.Id, err
	}

	return uid.String(), b.Id, nil
}

// GetCommunityBotList get community bot list.
func (svc *BotService) GetCommunityBotList(cid string, params *requestparams.ListPageParams) (*responseparams.ListPageResponse, error) {
	resp := &responseparams.ListPageResponse{}
	botInfos := make([]*responseparams.BotInfoResponse, 0)

	csvc := NewCommunityService((svc.ctx).(*gin.Context))
	info, err := csvc.GetCommunityInfoByUUIDV2(cid)
	if err != nil {
		return resp, err
	}

	// Get bot list
	bots, total, err := svc.SysBotDao.FindByCommunityIdPage(info.Id, (params.Page-1)*params.PageSize, params.PageSize)
	if err != nil {
		return resp, err
	}

	// Fill response data
	for _, b := range bots {
		botInfos = append(botInfos, &responseparams.BotInfoResponse{
			UUID:        b.UUID.String(),
			BotAddr:     b.BotAddr,
			BotPort:     b.BotPort,
			BotToken:    b.BotToken,
			Nickname:    b.Nickname,
			Avatar:      b.Avatar,
			Description: b.Description,
		})
	}

	resp.List = botInfos
	resp.Page = params.Page
	resp.PageSize = params.PageSize
	resp.Total = int(total)

	return resp, nil
}

// DeleteBot delete a bot.
func (svc *BotService) DeleteBot(uid string) error {
	if uid == "" {
		return common.NewError(common.I_INVALID_PARAM)
	}

	_uuid, err := uuid.FromString(uid)
	if err != nil {
		return err
	}

	u, err := svc.SysBotDao.FindFirstByUUID(_uuid)
	if err != nil {
		return err
	}

	err = svc.SysBotDao.Delete(u)
	if err != nil {
		return err
	}

	return nil
}

// UpdateBotInfo update bot info.
func (svc *BotService) UpdateBotInfo(params *requestparams.UpdateBotParams) (*responseparams.BotInfoResponse, error) {
	resp := &responseparams.BotInfoResponse{}

	if params.UUID == "" {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	_uuid, err := uuid.FromString(params.UUID)
	if err != nil {
		return resp, err
	}

	b, err := svc.SysBotDao.FindFirstByUUID(_uuid)
	if err != nil {
		return resp, err
	}

	// Update bot info
	if params.BotAddr != nil {
		b.BotAddr = *params.BotAddr
	}
	if params.BotPort != nil {
		b.BotPort = *params.BotPort
	}
	if params.BotToken != nil {
		b.BotToken = *params.BotToken
	}
	if params.Nickname != nil {
		b.Nickname = *params.Nickname
	}
	if params.Avatar != nil {
		b.Avatar = *params.Avatar
	}
	if params.Description != nil {
		b.Description = *params.Description
	}

	// Update user info
	if err = svc.SysBotDao.Update(b); err != nil {
		return resp, err
	}

	// Get new info
	b, _ = svc.SysBotDao.FindFirstByUUID(_uuid)
	resp.UUID = b.UUID.String()
	resp.BotAddr = b.BotAddr
	resp.BotPort = b.BotPort
	resp.BotToken = b.BotToken
	resp.Nickname = b.Nickname
	resp.Avatar = b.Avatar
	resp.Description = b.Description

	return resp, nil
}
