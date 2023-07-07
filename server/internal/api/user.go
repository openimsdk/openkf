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
	"github.com/OpenIMSDK/OpenKF/server/internal/param"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// AdminRegister
// @Tags user
// @Summary AdminRegister
// @Description Create a new admin
// @Produce application/json
// @Param data body param.RegisterAdminParams true "RegisterAdminParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/register/admin [post].
func AdminRegister(c *gin.Context) {
	var params param.RegisterAdminParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	_, _, err = svc.CreateAdmin(params)
	if err != nil {
		log.Debug("AdminRegister error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// StaffRegister
// @Tags user
// @Summary StaffRegister
// @Description Create a new staff
// @Produce application/json
// @Param data body param.RegisterStaffParams true "RegisterStaffParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/register/staff [post].
func StaffRegister(c *gin.Context) {
	var params param.RegisterStaffParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserService(c)
	_, _, err = svc.CreateStaff(params)
	if err != nil {
		log.Debug("AdminRegister error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}
