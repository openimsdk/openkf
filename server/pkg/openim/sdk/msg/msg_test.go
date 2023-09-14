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

package msg

import (
	"fmt"
	"testing"

	"github.com/openimsdk/openkf/server/pkg/openim/param/request"
	"github.com/openimsdk/openkf/server/pkg/openim/sdk/constant"
)

// TestAdminSendMsg test admin send msg function
func TestAdminSendMsg(t *testing.T) {
	adminToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOiJvcGVuSU1BZG1pbiIsIlBsYXRmb3JtSUQiOjEsImV4cCI6MTY5OTQyNTExMywibmJmIjoxNjkxNjQ4ODEzLCJpYXQiOjE2OTE2NDkxMTN9.JYNly2XpKkZuwZ9t4YTGki2TFst29fQYlmIFQytsoho"

	// test case
	testData := []struct {
		sendID           string
		recvID           string
		SenderPlatformID int
		content          string
		contentType      int
		sessionType      int
	}{
		{
			sendID:           "555248a1d1409a7abb5830fdad5d",
			recvID:           "54c09e9a6645bad1c6657dcee887",
			SenderPlatformID: constant.PLATFORMID_WEB,
			content:          "{\"content\":\"hello world!\"}",
			contentType:      constant.CONTENT_TYPE_TEXT,
			sessionType:      constant.SESSION_TYPE_SINGLE_CHAT,
		},
	}

	// range test case
	for _, data := range testData {
		res, err := AdminSendMsg(&request.MsgInfo{
			SendID:           data.sendID,
			RecvID:           data.recvID,
			SenderPlatformID: data.SenderPlatformID,
			Content:          &request.TextContent{Text: data.content},
			ContentType:      data.contentType,
			SessionType:      data.sessionType,
		},
			"123123123123123",
			"http://127.0.0.1:10002",
			adminToken,
		)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("%v", res)
	}
}
