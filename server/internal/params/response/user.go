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

package responseparams

// TokenResponse im token response.
type TokenResponse struct {
	Token             string `json:"token"`
	ExpireTimeSeconds uint   `json:"expire_time_seconds"`
}

// UserTokenResponse user token response.
type UserTokenResponse struct {
	UUID    string         `json:"uuid"`
	KFToken *TokenResponse `json:"kf_token"`
	IMToken *TokenResponse `json:"im_token"`
}

// UserInfoResponse user info response.
type UserInfoResponse struct {
	UUID        string `json:"uuid"`
	Email       string `json:"email"`
	Nickname    string `json:"nickname"`
	Avatar      string `json:"avatar"`
	Description string `json:"description"`
	IsEnable    bool   `json:"is_enable"`
	IsAdmin     bool   `json:"is_admin"`
	CreatedAt   string `json:"created_at"`
}

// UserStatisticItem user statistic item.
type UserStatisticItem struct {
	Data      interface{} `json:"data"`
	Timestamp int64       `json:"timestamp"`
}
