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

package response

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/openimsdk/openkf/server/internal/common"
)

// CommonCallbackResp Common callback response.
type CommonCallbackResp struct {
	ActionCode int    `json:"actionCode"`
	ErrCode    int32  `json:"errCode"`
	ErrMsg     string `json:"errMsg"`
	ErrDlt     string `json:"errDlt"`
}

// NewCommonCallbackResp Create a new CommonCallbackResp.
func NewCommonCallbackResp(actionCode int, errCode int32, errMsg string, errDlt string, c *gin.Context) {
	c.JSON(http.StatusOK, &CommonCallbackResp{
		ActionCode: actionCode,
		ErrCode:    errCode,
		ErrMsg:     errMsg,
		ErrDlt:     errDlt,
	})
}

// CallbackBeforeSendSingleMsgResp Response body for BeforeSendSingleMsgCommand.
type CallbackBeforeSendSingleMsgResp struct {
	CommonCallbackResp
}

// CallbackBeforeSendSingleMsgRespSuccess CallbackBeforeSendSingleMsgResp success response.
func CallbackBeforeSendSingleMsgRespSuccess(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackBeforeSendSingleMsgRespFail CallbackBeforeSendSingleMsgResp fail response.
func CallbackBeforeSendSingleMsgRespFail(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_DENY_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackAfterSendSingleMsgResp Response body for AfterSendSingleMsgCommand.
type CallbackAfterSendSingleMsgResp struct {
	CommonCallbackResp
}

// CallbackAfterSendSingleMsgRespSuccess CallbackAfterSendSingleMsgResp success response.
func CallbackAfterSendSingleMsgRespSuccess(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackAfterSendSingleMsgRespFail CallbackAfterSendSingleMsgResp fail response.
func CallbackAfterSendSingleMsgRespFail(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_DENY_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackMsgModifyCommandResp Response body for MsgModifyCommand.
type CallbackMsgModifyCommandResp struct {
	// todo: need to modify
	CommonCallbackResp
	Content          *string          `json:"content"`
	RecvID           *string          `json:"recvID"`
	GroupID          *string          `json:"groupID"`
	ClientMsgID      *string          `json:"clientMsgID"`
	ServerMsgID      *string          `json:"serverMsgID"`
	SenderPlatformID *int32           `json:"senderPlatformID"`
	SenderNickname   *string          `json:"senderNickname"`
	SenderFaceURL    *string          `json:"senderFaceURL"`
	SessionType      *int32           `json:"sessionType"`
	MsgFrom          *int32           `json:"msgFrom"`
	ContentType      *int32           `json:"contentType"`
	Status           *int32           `json:"status"`
	Options          *map[string]bool `json:"options"`
	OfflinePushInfo  interface{}      `json:"offlinePushInfo"`
	// OfflinePushInfo  *sdkws.OfflinePushInfo `json:"offlinePushInfo"`
	AtUserIDList *[]string `json:"atUserIDList"`
	MsgDataList  *[]byte   `json:"msgDataList"`
	AttachedInfo *string   `json:"attachedInfo"`
	Ex           *string   `json:"ex"`
}

// CallbackUserOnlineResp Response body for UserOnlineCommand.
type CallbackUserOnlineResp struct {
	CommonCallbackResp
}

// CallbackUserOnlineRespSuccess CallbackUserOnlineResp success response.
func CallbackUserOnlineRespSuccess(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackUserOnlineRespFail CallbackUserOnlineResp fail response.
func CallbackUserOnlineRespFail(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_DENY_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackUserOfflineResp Response body for UserOfflineCommand.
type CallbackUserOfflineResp struct {
	CommonCallbackResp
}

// CallbackUserOfflineRespSuccess CallbackUserOfflineResp success response.
func CallbackUserOfflineRespSuccess(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}

// CallbackUserOfflineRespFail CallbackUserOfflineResp fail response.
func CallbackUserOfflineRespFail(c *gin.Context) {
	NewCommonCallbackResp(
		common.OPENIM_SERVER_DENY_ACTION,
		common.OPENIM_SERVER_ALLOW_ACTION,
		common.GetMsg(common.OPENIM_SERVER_ALLOW_ACTION),
		"", // todo
		c,
	)
}
