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

	"github.com/openimsdk/openkf/server/internal/common"
	"github.com/openimsdk/openkf/server/internal/common/response"
	requestparams "github.com/openimsdk/openkf/server/internal/params/request"
	"github.com/openimsdk/openkf/server/internal/service"
	"github.com/openimsdk/openkf/server/internal/utils"
	"github.com/openimsdk/openkf/server/pkg/log"
)

// CreateCommunity
// @Tags community
// @Summary CreateCommunity
// @Description Create a new community
// @Produce application/json
// @Param data body requestparams.CommunityParams true "Community"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/community/create [post].
func CreateCommunity(c *gin.Context) {
	var params requestparams.CommunityParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewCommunityService(c)
	_, _, err = svc.Create(&params)
	if err != nil {
		log.Debug("CreateCommunity error", err)
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.Success(c)
}

// GetCommunityInfo
// @Tags community
// @Summary GetCommunityInfo
// @Description get community info
// @Produce application/json
// @Param data body requestparams.GetCommunityInfoParams true "GetCommunityInfoParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/community/info [post].
func GetCommunityInfo(c *gin.Context) {
	param := requestparams.GetCommunityInfoParams{}
	err := c.ShouldBindJSON(&param)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewCommunityService(c)
	resp, err := svc.GetCommunityInfoByUUID(param.UUID)
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// GetMyCommunityInfo
// @Tags community
// @Summary GetMyCommunityInfo
// @Description get my community info
// @Produce application/json
// @Security  ApiKeyAuth
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/community/me [get].
func GetMyCommunityInfo(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	svc := service.NewCommunityService(c)
	resp, err := svc.GetCommunityInfoByUUID(uuid)
	if err != nil {
		response.FailWithCode(common.KF_RECORD_NOT_FOUND, c)

		return
	}

	response.SuccessWithData(resp, c)
}

// GetMyCommunityInfo
// @Tags community
// @Summary GetMyCommunityInfo
// @Description get my community info
// @Produce application/json
// @Security  ApiKeyAuth
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/community/update [post].
func UpdateCommunity(c *gin.Context) {
	uuid, err := utils.GetCommunityUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	var params requestparams.UpdateCommunityInfoParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewCommunityService(c)
	info, err := svc.UpdateCommunity(uuid, &params)
	if err != nil {
		response.FailWithCode(common.ERROR, c)

		return
	}

	response.SuccessWithData(info, c)
}
