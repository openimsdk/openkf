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

// CreateStaff
// @Tags user
// @Summary CreateStaff
// @Description Create a new staff
// @Produce application/json
// @Security  ApiKeyAuth
// @Param data body requestparams.RegisterStaffParams true "RegisterStaffParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/staff/create [post].
func CreateStaff(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	var params requestparams.RegisterStaffParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	_, _, err = svc.CreateStaff(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// DeleteStaff
// @Tags user
// @Summary DeleteStaff
// @Description Delete a new staff
// @Produce application/json
// @Param data body requestparams.DeleteUserParams true "DeleteUserParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/staff/delete [post].
func DeleteStaff(c *gin.Context) {
	var params requestparams.DeleteUserParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	err = svc.DeleteStaff(params.UUID)
	if err != nil {
		log.Debug("DeleteStaff error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}
