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

package auth_test

import (
	"testing"

	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/sdk/auth"
)

// TestGetUserToken test get user token function
func TestGetUserToken(t *testing.T) {
	// test case
	testData := []struct {
		secret     string
		platformID uint
		userID     string
	}{
		{
			secret:     "openIM123",
			platformID: 5,
			userID:     "openIMAdmin",
		},
	}

	// range test case
	for _, data := range testData {
		token, err := auth.GetUserToken(&request.UserTokenParams{
			Secret:     data.secret,
			PlatformID: data.platformID,
			UserID:     data.userID,
		},
			"123123123123",
			"http://127.0.0.1:10002")
		if err != nil {
			t.Error(err)
		}
		t.Error(token)
	}
}
