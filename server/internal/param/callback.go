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

package param

// OpenIMCallbackCommand OpenIM callback request command.
type OpenIMCallbackCommand struct {
	Command string `form:"command" json:"command"`
}

// CommonCallbackReq Common callback request command.
type CommonCallbackReq struct {
	SendID           string   `json:"sendID"`
	CallbackCommand  string   `json:"callbackCommand"`
	ServerMsgID      string   `json:"serverMsgID"`
	ClientMsgID      string   `json:"clientMsgID"`
	OperationID      string   `json:"operationID"`
	SenderPlatformID int32    `json:"senderPlatformID"`
	SenderNickname   string   `json:"senderNickname"`
	SessionType      int32    `json:"sessionType"`
	MsgFrom          int32    `json:"msgFrom"`
	ContentType      int32    `json:"contentType"`
	Status           int32    `json:"status"`
	CreateTime       int64    `json:"createTime"`
	Content          string   `json:"content"`
	Seq              uint32   `json:"seq"`
	AtUserIDList     []string `json:"atUserList"`
	SenderFaceURL    string   `json:"faceURL"`
	Ex               string   `json:"ex"`
}

// CallbackBeforeSendSingleMsgReq Request body for BeforeSendSingleMsgCommand.
type CallbackBeforeSendSingleMsgReq struct {
	CommonCallbackReq
	RecvID string `json:"recvID"`
}

// CallbackAfterSendSingleMsgReq Request body for AfterSendSingleMsgCommand.
type CallbackAfterSendSingleMsgReq struct {
	CommonCallbackReq
	RecvID string `json:"recvID"`
}

// CallbackMsgModifyCommandReq Request body for MsgModifyCommand.
type CallbackMsgModifyCommandReq struct {
	CommonCallbackReq
}

// UserStatusBaseCallback Common user callback request command.
type UserStatusBaseCallback struct {
	CallbackCommand string `json:"callbackCommand"`
	OperationID     string `json:"operationID"`
	PlatformID      int    `json:"platformID"`
	Platform        string `json:"platform"`
	UserID          string `json:"userID"`
}

// CallbackUserOnlineReq Request body for UserOnlineCommand.
type CallbackUserOnlineReq struct {
	UserStatusBaseCallback
	// Token           string `json:"token"`
	Seq             int64  `json:"seq"`
	IsAppBackground bool   `json:"isAppBackground"`
	ConnID          string `json:"connID"`
}

// CallbackUserOfflineReq Request body for UserOfflineCommand.
type CallbackUserOfflineReq struct {
	UserStatusBaseCallback
	Seq    int64  `json:"seq"`
	ConnID string `json:"connID"`
}
