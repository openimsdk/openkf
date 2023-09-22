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
	responseparams "github.com/openimsdk/openkf/server/internal/params/response"
	"github.com/openimsdk/openkf/server/internal/service"
	"github.com/openimsdk/openkf/server/internal/utils"
)

// GetUserStatistics
// @Tags user
// @Summary GetUserStatistics
// @Description get user statistics
// @Security  ApiKeyAuth
// @Produce application/json
// @Param data body requestparams.GetUsersStatisticsParams true "GetUsersStatisticsParams"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/user/statistics [post].
func GetUserStatistics(c *gin.Context) {
	uuid, err := utils.GetUserUUID(c)
	if err != nil {
		response.FailWithCode(common.UNAUTHORIZED, c)

		return
	}

	var params requestparams.GetUsersStatisticsParams
	err = c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	svc := service.NewUserStatisticService(c)
	var result []*responseparams.UserStatisticItem

	switch params.Type {
	case string(requestparams.USER_STATISTICS_ONLINE_TIME):
		result, err = svc.ReadUserOnlineTimePerDay(uuid, params.StartTimestamp, params.EndTimestamp)
		if err != nil {
			response.FailWithData(err.Error(), c)

			return
		}
	case string(requestparams.USER_STATISTICS_MESSAGE_COUNT):
		result, err = svc.ReadUserSendMsgCountPerDay(uuid, params.StartTimestamp, params.EndTimestamp)
		if err != nil {
			response.FailWithData(err.Error(), c)

			return
		}
	case string(requestparams.USER_STATISTICS_MESSAGE_LENGTH):
		result, err = svc.ReadUserSendMsgLenPerDay(uuid, params.StartTimestamp, params.EndTimestamp)
		if err != nil {
			response.FailWithData(err.Error(), c)

			return
		}
	}

	response.SuccessWithData(result, c)
}
