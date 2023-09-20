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
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/openimsdk/openkf/server/internal/common"
	"github.com/openimsdk/openkf/server/internal/common/response"
	slackcmd "github.com/openimsdk/openkf/server/internal/msg/slack_cmd"
	requestparams "github.com/openimsdk/openkf/server/internal/params/request"
	"github.com/openimsdk/openkf/server/internal/service"
	"github.com/openimsdk/openkf/server/pkg/log"
)

// OpenIMCallback
// @Tags openim
// @Summary OpenIMCallback
// @Description Support OpenIM callback
// @Produce application/json
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/openim/callback [post].
func OpenIMCallback(c *gin.Context) {
	param := &requestparams.OpenIMCallbackCommand{}
	err := c.ShouldBindQuery(param)
	if err != nil {
		// todo
		response.FailWithCode(common.INVALID_PARAMS, c)

		return
	}

	// redirect to the corresponding router
	callbackURL := fmt.Sprintf("%s%s", c.Request.URL.Path, param.Command)
	c.Request.URL.Path = callbackURL
	c.Redirect(http.StatusTemporaryRedirect, callbackURL)
}

// BeforeSendSingleMsg.
func BeforeSendSingleMsg(c *gin.Context) {
	// Msg filter
	params := &requestparams.MsgAbstract{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		// Do not check
		return
	}

	// Send message to slack user
	ssvc := service.NewUserDispatchService(c)
	if check := ssvc.SlackUserFilter(params.RecvID); check {
		log.Debugf("Msg filter", "Send message to slack: %s", params.RecvID)
		slackcmd.SendMsg(params)
	}
}

// AfterSendSingleMsg.
func AfterSendSingleMsg(c *gin.Context) {
	log.Infof("UserOnline", "UserOnline")
	params := &requestparams.MsgAbstract{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		// Do not check
		return
	}

	// Save to influxdb
	svc := service.NewIMCallbackService(c)
	if err := svc.AfterSendSingleMsgCallback(params.SendID, params.Content); err != nil {
		log.Errorf("AfterSendSingleMsg", "Write to influx db error: %s", err.Error())

		return
	}
}

// UserOnline.
func UserOnline(c *gin.Context) {
	log.Infof("UserOnline", "UserOnline")
	params := &requestparams.CallbackUserOnlineReq{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		// Do not check
		return
	}

	// Save to influxdb
	svc := service.NewIMCallbackService(c)
	if err := svc.UserOnlineCallback(params.UserID); err != nil {
		log.Errorf("UserOnline", "Write to influx db error: %s", err.Error())

		return
	}
}

// UserOffline.
func UserOffline(c *gin.Context) {
	log.Infof("UserOffline", "UserOffline")
	params := &requestparams.CallbackUserOfflineReq{}
	err := c.ShouldBindJSON(params)
	if err != nil {
		// Do not check
		return
	}

	// Save to influxdb
	svc := service.NewIMCallbackService(c)
	if err := svc.UserOfflineCallback(params.UserID); err != nil {
		log.Errorf("UserOffline", "Write to influx db error: %s", err.Error())

		return
	}
}

// MsgModify.
func MsgModify(c *gin.Context) {
}

// OfflinePush.
func OfflinePush(c *gin.Context) {
}

// OnlinePush.
func OnlinePush(c *gin.Context) {
}
