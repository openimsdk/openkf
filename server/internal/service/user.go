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
	"fmt"
	"net"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	"github.com/OpenIMSDK/OpenKF/server/internal/param"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/user"
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

	// TODO: set pipline to tx.
	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   uuid.String(),
				Nickname: adminParam.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.SysUserDao.Delete(u)

		return uuid.String(), u.Id, err
	}

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

	// TODO: set pipline to tx.
	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   uuid.String(),
				Nickname: staffParam.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.SysUserDao.Delete(u)

		return uuid.String(), u.Id, err
	}

	return uuid.String(), u.Id, nil
}

// registerUserToOpenIM register user to openim.
func registerUserToOpenIM(param *request.RegisterUserParams) (bool, error) {
	// Default not use tls/ssl
	host := fmt.Sprintf("http://%s", net.JoinHostPort(config.Config.OpenIM.Ip, fmt.Sprintf("%d", config.Config.OpenIM.ApiPort)))
	resp, err := user.RegisterUser(param, host)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New(resp.ErrMsg)
	}

	return true, nil
}
