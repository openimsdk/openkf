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
)

// SlackConfig
// @Tags platform
// @Summary SlackConfig
// @Description Get slack config
// @Produce application/json
// @Security  ApiKeyAuth
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/platform/slack/config [get].
func SlackConfig(c *gin.Context) {
	// TODO: Check user role.
	if _, err := utils.GetUserUUID(c); err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewSlackService(c)
	// TODO: Add a table and save config to db.
	// Now it only read from config file.
	resp, err := svc.GetSlackConfig()
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// GetSlackCustomer
// @Tags platform
// @Summary GetSlackCustomer
// @Description get slack user info
// @Security  ApiKeyAuth
// Param data body param.GetUserInfoParams true "GetUserInfoParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/platform/slack/customer [post].
func GetSlackCustomer(c *gin.Context) {
	// TODO: Check role.

	param := requestparams.GetUserInfoParams{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewSlackService(c)
	resp, err := svc.GetSlackUser(param.UUID)
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}
