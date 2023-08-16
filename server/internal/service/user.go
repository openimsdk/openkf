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

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/config"
	"github.com/OpenIMSDK/OpenKF/server/internal/dal/dao"
	"github.com/OpenIMSDK/OpenKF/server/internal/models/base"
	systemroles "github.com/OpenIMSDK/OpenKF/server/internal/models/system_roles"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	responseparams "github.com/OpenIMSDK/OpenKF/server/internal/params/response"
	internal_utils "github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/response"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/auth"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/user"
	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// UserService user service.
type UserService struct {
	Service

	SysUserDao      *dao.SysUserDao
	SysCommunityDao *dao.SysCommunityDao
}

// NewUserService return new service with gin context.
func NewUserService(c *gin.Context) *UserService {
	return &UserService{
		Service: Service{
			ctx: c,
		},
		SysUserDao:      dao.NewSysUserDao(),
		SysCommunityDao: dao.NewSysCommunityDao(),
	}
}

// CreateAdmin create admin user.
func (svc *UserService) CreateAdmin(user *requestparams.RegisterAdminParams) (string, uint, error) {
	// Check code
	mService := NewMailService((svc.ctx).(*gin.Context))
	if isExist := mService.CheckCode(user.UserInfo.Email, user.Code); !isExist {
		return "", 0, errors.New("code is not valid")
	}

	// Create community
	communityParam := user.CommunityInfo
	cService := NewCommunityService((svc.ctx).(*gin.Context))
	_, cid, err := cService.Create(&communityParam)
	if err != nil {
		return "", 0, err
	}

	// Create admin
	uuid := utils.GenUUIDWithoutHyphen()
	adminParam := user.UserInfo
	admin := &systemroles.SysUser{
		UserBase: base.UserBase{
			UUID:     uuid,
			Email:    adminParam.Email,
			Nickname: adminParam.Nickname,
			Avatar:   *adminParam.Avatar,
			IsEnable: true,
		},
		IsAdmin:     true,
		Password:    utils.EncryptPassword(adminParam.Password),
		CommunityId: cid,
	}
	if err = svc.SysUserDao.Create(admin); err != nil {
		return uuid, 0, err
	}

	u, _ := svc.SysUserDao.FindFirstByUUID(uuid)

	// TODO: set pipline to tx.
	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   uuid,
				Nickname: adminParam.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.SysUserDao.Delete(u)

		return uuid, u.Id, err
	}

	return uuid, u.Id, nil
}

// CreateStaff create staff user.
func (svc *UserService) CreateStaff(cid string, user *requestparams.RegisterStaffParams) (string, uint, error) {
	// Create staff
	uid := utils.GenUUIDWithoutHyphen()

	communityInfo, err := svc.SysCommunityDao.FindFirstByUUID(cid)
	if err != nil {
		return "", 0, err
	}

	staffParam := user.UserInfo
	staff := &systemroles.SysUser{
		UserBase: base.UserBase{
			UUID:     uid,
			Email:    staffParam.Email,
			Nickname: staffParam.Nickname,
			Avatar:   *staffParam.Avatar,
			IsEnable: true,
		},
		IsAdmin:     false,
		Password:    utils.EncryptPassword(staffParam.Password),
		CommunityId: communityInfo.Id,
	}
	if err = svc.SysUserDao.Create(staff); err != nil {
		return uid, 0, err
	}

	// TODO: Send email to staff

	u, _ := svc.SysUserDao.FindFirstByUUID(uid)

	// TODO: set pipline to tx.
	param := &request.RegisterUserParams{
		Secret: config.Config.OpenIM.Secret,
		Users: []request.User{
			{
				UserID:   uid,
				Nickname: staffParam.Nickname,
				FaceURL:  "", // Use OpenKF avatar
			},
		},
	}
	ok, err := registerUserToOpenIM(param)
	if err != nil || !ok {
		// Assume that the user has been created/deleted successfully
		_ = svc.SysUserDao.Delete(u)

		return uid, u.Id, err
	}

	return uid, u.Id, nil
}

// registerUserToOpenIM register user to openim.
func registerUserToOpenIM(param *request.RegisterUserParams) (bool, error) {
	// TODO: Add get operationID

	// Default not use tls/ssl
	host := fmt.Sprintf("http://%s", net.JoinHostPort(config.Config.OpenIM.Ip, fmt.Sprintf("%d", config.Config.OpenIM.ApiPort)))
	resp, err := user.RegisterUser(param, "registerUserToOpenIM", host)
	if err != nil {
		return false, err
	}

	if resp.ErrCode != 0 {
		return false, errors.New(resp.ErrMsg)
	}

	return true, nil
}

// DeleteStaff delete staff user.
func (svc *UserService) DeleteStaff(uid string) error {
	if uid == "" {
		return common.NewError(common.I_INVALID_PARAM)
	}

	u, err := svc.SysUserDao.FindFirstByUUID(uid)
	if err != nil {
		return err
	}

	err = svc.SysUserDao.Delete(u)
	if err != nil {
		return err
	}

	return nil
}

// LoginWithAccount login with account.
func (svc *UserService) LoginWithAccount(param *requestparams.LoginParamsWithAccount) (*responseparams.UserTokenResponse, error) {
	resp := &responseparams.UserTokenResponse{}

	// Check user
	u, err := svc.SysUserDao.FindFirstByEmail(param.Email)
	if err != nil {
		return resp, err
	}

	// Check password
	if !utils.ComparePassword(param.Password, u.Password) {
		return resp, errors.New("password is not correct")
	}

	// Get community id
	cService := NewCommunityService((svc.ctx).(*gin.Context))
	c, err := cService.GetCommunityInfoById(u.CommunityId)
	if err != nil {
		return resp, err
	}

	// Generate KF token
	kfToken, kfExpireTimeSeconds, err := internal_utils.GenerateJwtToken(u.UUID, c.UUID)
	if err != nil {
		return resp, err
	}

	// Get IM token
	imParam := &request.UserTokenParams{
		Secret:     config.Config.OpenIM.Secret,
		UserID:     u.UUID,
		PlatformID: uint(config.Config.OpenIM.PlatformID),
	}
	imResp, err := getUserIMToken(imParam)
	if err != nil {
		return resp, err
	}

	// Fill response data
	resp.UUID = u.UUID
	resp.KFToken = &responseparams.TokenResponse{
		Token:             kfToken,
		ExpireTimeSeconds: kfExpireTimeSeconds,
	}
	resp.IMToken = &responseparams.TokenResponse{
		Token:             imResp.Token,
		ExpireTimeSeconds: imResp.ExpireTimeSeconds,
	}

	// TODO: Set Online in OpenIM or do this in js-sdk

	return resp, nil
}

// getUserIMToken get user im token.
func getUserIMToken(param *request.UserTokenParams) (*response.TokenData, error) {
	// TODO: Add operationID

	// Default not use tls/ssl
	host := fmt.Sprintf("http://%s", net.JoinHostPort(config.Config.OpenIM.Ip, fmt.Sprintf("%d", config.Config.OpenIM.ApiPort)))
	resp, err := auth.GetUserToken(param, "getUserIMToken", host)
	if err != nil {
		return &response.TokenData{}, err
	}

	if resp.ErrCode != 0 {
		return &response.TokenData{}, errors.New(resp.ErrMsg)
	}

	return &resp.Data, nil
}

// GetAdminToken get admin token.
func (svc *UserService) GetAdminToken() (*response.TokenData, error) {
	params := &request.UserTokenParams{
		Secret:     config.Config.OpenIM.Secret,
		PlatformID: uint(config.Config.OpenIM.PlatformID),
		UserID:     config.Config.OpenIM.AdminID,
	}

	// TODO: Get cache from redis

	return getUserIMToken(params)
}

// GetUserInfoByUUID get user info by uuid.
func (svc *UserService) GetUserInfoByUUID(uid string) (*responseparams.UserInfoResponse, error) {
	resp := &responseparams.UserInfoResponse{}

	if uid == "" {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	u, err := svc.SysUserDao.FindFirstByUUID(uid)
	if err != nil {
		return resp, err
	}

	resp.UUID = u.UUID
	resp.Email = u.Email
	resp.Nickname = u.Nickname
	resp.Avatar = u.Avatar
	resp.Description = u.Description
	resp.IsAdmin = u.IsAdmin
	resp.IsEnable = u.IsEnable
	resp.CreatedAt = u.CreatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

// UpdateUserInfo update user info.
func (svc *UserService) UpdateUserInfo(uid string, params *requestparams.UpdateUserInfoParams) (*responseparams.UserInfoResponse, error) {
	resp := &responseparams.UserInfoResponse{}

	if uid == "" {
		return resp, common.NewError(common.I_INVALID_PARAM)
	}

	u, err := svc.SysUserDao.FindFirstByUUID(uid)
	if err != nil {
		return resp, err
	}

	// Update user info
	if params.Nickname != nil {
		u.Nickname = *params.Nickname
	}
	if params.Description != nil {
		u.Description = *params.Description
	}
	if params.Avatar != nil {
		u.Avatar = *params.Avatar
	}
	if params.Email != nil {
		u.Email = *params.Email
	}

	// TODO: Check if the user is admin
	if params.IsEnable != nil {
		u.IsEnable = *params.IsEnable
	}
	if params.IsAdmin != nil {
		u.IsAdmin = *params.IsAdmin
	}

	// Update user info
	if err = svc.SysUserDao.Update(u); err != nil {
		return resp, err
	}

	// Get new info
	u, _ = svc.SysUserDao.FindFirstByUUID(uid)
	resp.UUID = u.UUID
	resp.Email = u.Email
	resp.Nickname = u.Nickname
	resp.Avatar = u.Avatar
	resp.Description = u.Description
	resp.IsAdmin = u.IsAdmin
	resp.IsEnable = u.IsEnable
	resp.CreatedAt = u.CreatedAt.Format("2006-01-02 15:04:05")

	return resp, nil
}

// UpdateUserPassword update user password.
func (svc *UserService) UpdateUserPassword(uid string, params *requestparams.UpdateUserPasswordParams) error {
	if uid == "" {
		return common.NewError(common.I_INVALID_PARAM)
	}

	u, err := svc.SysUserDao.FindFirstByUUID(uid)
	if err != nil {
		return err
	}

	// Update user info
	if params.Password != "" && params.RepeatPassword != "" &&
		params.Password == params.RepeatPassword {
		u.Password = utils.EncryptPassword(params.Password)
	}

	// Update user info
	if err = svc.SysUserDao.Update(u); err != nil {
		return err
	}

	return nil
}

// GetCommunityUserList get community user list.
func (svc *UserService) GetCommunityUserList(cid string, params *requestparams.ListPageParams) (*responseparams.ListPageResponse, error) {
	resp := &responseparams.ListPageResponse{}
	userInfos := make([]*responseparams.UserInfoResponse, 0)

	csvc := NewCommunityService((svc.ctx).(*gin.Context))
	info, err := csvc.GetCommunityInfoByUUIDV2(cid)
	if err != nil {
		return resp, err
	}

	// Get user list
	users, total, err := svc.SysUserDao.FindByCommunityIdPage(info.Id, (params.Page-1)*params.PageSize, params.PageSize)
	if err != nil {
		return resp, err
	}

	// Fill response data
	for _, u := range users {
		userInfos = append(userInfos, &responseparams.UserInfoResponse{
			UUID:        u.UUID,
			Email:       u.Email,
			Nickname:    u.Nickname,
			Avatar:      u.Avatar,
			Description: u.Description,
			IsAdmin:     u.IsAdmin,
			IsEnable:    u.IsEnable,
			CreatedAt:   u.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	resp.List = userInfos
	resp.Page = params.Page
	resp.PageSize = params.PageSize
	resp.Total = int(total)

	return resp, nil
}
