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

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/db"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	"github.com/OpenIMSDK/OpenKF/server/internal/param"
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
func (svc *CommunityService) Create(community param.CommunityParams) (string, uint, error) {
	uuid := utils.GenUUID()
	data := &systemroles.SysCommunity{
		UUID:   uuid,
		Name:   community.Name,
		Email:  community.Email,
		Avatar: community.Avatar,
	}

	err := db.GetMysqlDB().Create(data).Error
	// err := svc.SysCommunityDao.Create(data)
	if err != nil {
		return uuid.String(), 0, err
	}

	// Get community id
	c, _ := svc.SysCommunityDao.FindFirstByUUID(uuid)

	return c.UUID.String(), c.Id, nil
}

// GetCommunityInfoByUUID get community info by uuid.
func (svc *CommunityService) GetCommunityInfoByUUID(uid string, offset int, limit int) ([]*systemroles.SysCommunity, int64, error) {
	_uuid := uuid.Must(uuid.FromString(uid))
	c, count, err := svc.SysCommunityDao.FindByUUIDPage(_uuid, offset, limit)

	return c, count, err
}
