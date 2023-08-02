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

// CreateBot
// @Tags bot
// @Summary CreateBot
// @Description Create a new bot
// @Produce application/json
// @Security  ApiKeyAuth
// @Param data body requestparams.CreateBotParams true "CreateBotParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/bot/create [post].
func CreateBot(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	var params requestparams.CreateBotParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewBotService(c)
	_, _, err = svc.CreateBot(uuid, &params)
	if err != nil {
		response.FailWithAll(common.ERROR, err.Error(), c)

		return
	}

	response.Success(c)
}

// ListBot
// @Tags bot
// @Summary ListBot
// @Description get community bot info
// @Security  ApiKeyAuth
// Param data body param.ListPageParams true "ListPageParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/bot/list [post].
func ListBot(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewBotService(c)
	var params requestparams.ListPageParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	resp, err := svc.GetCommunityBotList(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// DeleteBot
// @Tags bot
// @Summary DeleteBot
// @Description delete bot
// @Security  ApiKeyAuth
// Param data body param.DeleteBotParams true "DeleteBotParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/bot/delete [post].
func DeleteBot(c *gin.Context) {
	var params requestparams.DeleteBotParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewBotService(c)
	err = svc.DeleteBot(params.UUID)
	if err != nil {
		log.Debug("DeleteBot error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// UpdateBot
// @Tags bot
// @Summary UpdateBot
// @Description update bot info
// @Security  ApiKeyAuth
// Param data body param.UpdateBotParams true "UpdateBotParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/admin/bot/update [post].
func UpdateBot(c *gin.Context) {
	var params requestparams.UpdateBotParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewBotService(c)
	info, err := svc.UpdateBotInfo(&params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(info, c)
}
