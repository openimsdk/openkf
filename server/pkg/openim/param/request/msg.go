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

package request

// MsgInfo msg info.
type MsgInfo struct {
	SendID           string           `json:"sendID"           binding:"required"`
	RecvID           string           `json:"recvID"           binding:"required"`
	GroupID          string           `json:"groupID"`
	SenderNickname   string           `json:"senderNickname"`
	SenderFaceURL    string           `json:"senderFaceURL"`
	SenderPlatformID int              `json:"senderPlatformID"`
	Content          *TextContent     `json:"content"          binding:"required"`
	ContentType      int              `json:"contentType"      binding:"required"`
	SessionType      int              `json:"sessionType"      binding:"required"`
	IsOnlineOnly     bool             `json:"isOnlineOnly"`
	NotOfflinePush   bool             `json:"notOfflinePush"`
	OfflinePushInfo  *OfflinePushInfo `json:"offlinePushInfo"`
}

// TextContent text content.
type TextContent struct {
	Text string `json:"content"`
}

// OfflinePushInfo offline push info.
type OfflinePushInfo struct {
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Ex            string `json:"ex"`
	IOSPushSound  string `json:"iOSPushSound"`
	IOSBadgeCount bool   `json:"iOSBadgeCount"`
}
