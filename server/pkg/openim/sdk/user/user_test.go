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

package user_test

import (
	"testing"

	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/user"
)

// TestGetUserToken test get user token function
func TestRegisterUser(t *testing.T) {
	// test case
	testData := []struct {
		secret   string
		userID   string
		nickname string
		faceURL  string
	}{
		{
			secret:   "openkf",
			userID:   "test123",
			nickname: "test123",
			faceURL:  "https://openim.com/openkf.png",
		},
	}

	// range test case
	for _, data := range testData {
		_, err := user.RegisterUser(&request.RegisterUserParams{
			Secret: data.secret,
			Users: []request.User{
				{
					UserID:   data.userID,
					Nickname: data.nickname,
					FaceURL:  data.faceURL,
				},
			},
		},
			"123123123",
			"http://127.0.0.1:10002")
		if err != nil {
			t.Error(err)
		}
	}
}
