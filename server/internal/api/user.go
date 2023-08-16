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

package api

import (
	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/common/response"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// AdminRegister
// @Tags user
// @Summary AdminRegister
// @Description Create a new admin
// @Produce application/json
// @Param data body requestparams.RegisterAdminParams true "RegisterAdminParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/register/admin [post].
func AdminRegister(c *gin.Context) {
	var params requestparams.RegisterAdminParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	_, _, err = svc.CreateAdmin(&params)
	if err != nil {
		log.Debug("AdminRegister error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// AccountLogin
// @Tags user
// @Summary AccountLogin
// @Description login with account
// @Produce application/json
// @Param data body requestparams.LoginParamsWithAccount true "LoginParamsWithAccount"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/login/account [post].
func AccountLogin(c *gin.Context) {
	var params requestparams.LoginParamsWithAccount
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	u, err := svc.LoginWithAccount(&params)
	if err != nil {
		log.Debug("Login error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	// TODO: Add login status to user dispatch
	dsvc := service.NewUserDispatchService(c.Request.Context())
	if err = dsvc.AddUser(u.UUID); err != nil {
		log.Error("Add to online status error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(u, c)
}

// AccountExit
// @Tags user
// @Summary AccountExit
// @Description logout with account
// @Produce application/json
// @Security  ApiKeyAuth
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/login/account [post].
func AccountExit(c *gin.Context) {
	uuid, err := utils.GetUserUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewUserDispatchService(c)
	err = svc.DeleteUser(uuid)
	if err != nil {
		log.Error("Delete user from online status error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// GetUserInfo
// @Tags user
// @Summary GetUserInfo
// @Description get user info
// @Security  ApiKeyAuth
// Param data body param.GetUserInfoParams true "GetUserInfoParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/info [post].
func GetUserInfo(c *gin.Context) {
	param := requestparams.GetUserInfoParams{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	resp, err := svc.GetUserInfoByUUID(param.UUID)
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// GetMyInfo
// @Tags user
// @Summary GetMyInfo
// @Description get my user info
// @Security  ApiKeyAuth
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/me [get].
func GetMyInfo(c *gin.Context) {
	uuid, err := utils.GetUserUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewUserService(c)
	resp, err := svc.GetUserInfoByUUID(uuid)
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// GetCommunityUserList
// @Tags user
// @Summary GetCommunityUserList
// @Description get community user info
// @Security  ApiKeyAuth
// Param data body param.ListPageParams true "ListPageParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/userlist [post].
func GetCommunityUserList(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewUserService(c)
	var params requestparams.ListPageParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	resp, err := svc.GetCommunityUserList(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// UpdatePassword
// @Tags user
// @Summary UpdatePassword
// @Description update user password
// @Security  ApiKeyAuth
// Param data body param.UpdateUserPasswordParams true "UpdateUserPasswordParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/update-password [post].
func UpdatePassword(c *gin.Context) {
	uuid, err := utils.GetUserUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	var params requestparams.UpdateUserPasswordParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	err = svc.UpdateUserPassword(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// UpdateInfo
// @Tags user
// @Summary UpdateInfo
// @Description update user info
// @Security  ApiKeyAuth
// Param data body param.UpdateUserInfoParams true "UpdateUserInfoParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/update [post].
func UpdateInfo(c *gin.Context) {
	uuid, err := utils.GetUserUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	var params requestparams.UpdateUserInfoParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	info, err := svc.UpdateUserInfo(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(info, c)
}
