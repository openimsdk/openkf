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
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// CommunityService community service.
type CommunityService struct {
	Service

	SysCommunityDao *dao.SysCommunityDao
}

// NewCommunityService return new service with gin context.
func NewCommunityService(c *gin.Context) *CommunityService {
	return &CommunityService{
		Service: Service{
			ctx: c,
		},
		SysCommunityDao: dao.NewSysCommunityDao(),
	}
}

// Create create community.
func (svc *CommunityService) Create(community *requestparams.CommunityParams) (string, uint, error) {
	uuid := utils.GenUUID()
	data := &systemroles.SysCommunity{
		UUID:        uuid,
		Name:        community.Name,
		Email:       community.Email,
		Description: *community.Description,
		Avatar:      *community.Avatar,
	}

	// err := db.GetMysqlDB().Create(data).Error
	err := svc.SysCommunityDao.Create(data)
	if err != nil {
		return uuid.String(), 0, err
	}

	// Get community id
	c, _ := svc.SysCommunityDao.FindFirstByUUID(uuid)

	return c.UUID.String(), c.Id, nil
}

// GetCommunityInfoById get community info by id.
func (svc *CommunityService) GetCommunityInfoById(id uint) (*responseparams.CommunityInfoResponse, error) {
	resp := &responseparams.CommunityInfoResponse{}

	if id <= 0 {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	c, err := svc.SysCommunityDao.FindFirstById(id)
	if err != nil {
		return resp, err
	}

	resp.UUID = c.UUID.String()
	resp.Email = c.Email
	resp.Avatar = c.Avatar
	resp.Name = c.Name
	resp.Description = c.Description

	return resp, err
}

// GetCommunityInfoByUUID get community info by uuid.
func (svc *CommunityService) GetCommunityInfoByUUID(uid string) (*responseparams.CommunityInfoResponse, error) {
	resp := &responseparams.CommunityInfoResponse{}

	if uid == "" {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	_uuid, err := uuid.FromString(uid)
	if err != nil {
		return resp, err
	}

	c, err := svc.SysCommunityDao.FindFirstByUUID(_uuid)
	if err != nil {
		return resp, err
	}

	resp.UUID = c.UUID.String()
	resp.Email = c.Email
	resp.Avatar = c.Avatar
	resp.Name = c.Name
	resp.Description = c.Description

	return resp, nil
}

// GetCommunityInfoByUUIDV2 get community info by uuid and return SysCommunity.
func (svc *CommunityService) GetCommunityInfoByUUIDV2(uid string) (*systemroles.SysCommunity, error) {
	if uid == "" {
		return nil, common.NewError(common.I_INVALID_PARAM)
	}

	_uuid, err := uuid.FromString(uid)
	if err != nil {
		return nil, err
	}

	c, err := svc.SysCommunityDao.FindFirstByUUID(_uuid)
	if err != nil {
		return nil, err
	}

	return c, nil
}

// UpdateCommunity update community info.
func (svc *CommunityService) UpdateCommunity(uid string, community *requestparams.UpdateCommunityInfoParams) (*responseparams.CommunityInfoResponse, error) {
	resp := &responseparams.CommunityInfoResponse{}

	if uid == "" {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	_uuid, err := uuid.FromString(uid)
	if err != nil {
		return resp, err
	}

	info, err := svc.SysCommunityDao.FindFirstByUUID(_uuid)
	if err != nil {
		return resp, err
	}

	if community.Name != nil {
		info.Name = *community.Name
	}
	if community.Email != nil {
		info.Email = *community.Email
	}
	if community.Description != nil {
		info.Description = *community.Description
	}
	if community.Avatar != nil {
		info.Avatar = *community.Avatar
	}

	// Update community info
	err = svc.SysCommunityDao.Update(info)
	if err != nil {
		return resp, err
	}

	resp.UUID = info.UUID.String()
	resp.Email = info.Email
	resp.Avatar = info.Avatar
	resp.Name = info.Name
	resp.Description = info.Description

	return resp, nil
}
