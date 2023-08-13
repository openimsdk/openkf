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

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/common/response"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
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
}

// AfterSendSingleMsg.
func AfterSendSingleMsg(c *gin.Context) {
}

// MsgModify.
func MsgModify(c *gin.Context) {
}

// UserOnline.
func UserOnline(c *gin.Context) {
	// TODO: push user id to queue
}

// UserOffline.
func UserOffline(c *gin.Context) {
	// TODO: remove user id to queue
}

// OfflinePush.
func OfflinePush(c *gin.Context) {
}

// OnlinePush.
func OnlinePush(c *gin.Context) {
}
