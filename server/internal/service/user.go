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
	"errors"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	"github.com/OpenIMSDK/OpenKF/server/internal/param"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// UserService user service.
type UserService struct {
	Service

	SysUserDao *dao.SysUserDao
}

// NewUserService return new service with gin context.
func NewUserService(c *gin.Context) *UserService {
	return &UserService{
		Service: Service{
			ctx: c,
		},
		SysUserDao: dao.NewSysUserDao(),
	}
}

// CreateAdmin create admin user.
func (svc *UserService) CreateAdmin(user param.RegisterAdminParams) (string, uint, error) {
	// Check code
	mService := NewMailService((svc.ctx).(*gin.Context))
	if isExist := mService.CheckCode(user.UserInfo.Email, user.Code); !isExist {
		return "", 0, errors.New("code is not valid")
	}

	// Create community
	communityParam := user.CommunityInfo
	cService := NewCommunityService((svc.ctx).(*gin.Context))
	_, cid, err := cService.Create(communityParam)
	if err != nil {
		return "", 0, err
	}

	// Create admin
	uuid := utils.GenUUID()
	adminParam := user.UserInfo
	admin := &systemroles.SysUser{
		UserBase: base.UserBase{
			UUID:     uuid,
			Email:    adminParam.Email,
			Nickname: adminParam.Nickname,
			Avatar:   adminParam.Avatar,
			IsEnable: true,
		},
		IsAdmin:     true,
		Password:    utils.EncryptPassword(adminParam.Password),
		CommunityId: cid,
	}
	if err = svc.SysUserDao.Create(admin); err != nil {
		return uuid.String(), 0, err
	}

	u, _ := svc.SysUserDao.FindFirstByUUID(uuid)

	return uuid.String(), u.Id, nil
}

// CreateStaff create staff user.
func (svc *UserService) CreateStaff(user param.RegisterStaffParams) (string, uint, error) {
	// Create staff
	uuid := utils.GenUUID()
	staffParam := user.UserInfo
	staff := &systemroles.SysUser{
		UserBase: base.UserBase{
			UUID:     uuid,
			Email:    staffParam.Email,
			Nickname: staffParam.Nickname,
			Avatar:   staffParam.Avatar,
			IsEnable: true,
		},
		IsAdmin:     false,
		Password:    utils.EncryptPassword(staffParam.Password),
		CommunityId: user.CommunityId,
	}
	if err := svc.SysUserDao.Create(staff); err != nil {
		return uuid.String(), 0, err
	}

	// TODO: Send email to staff

	u, _ := svc.SysUserDao.FindFirstByUUID(uuid)

	return uuid.String(), u.Id, nil
}