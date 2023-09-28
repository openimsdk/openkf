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

// BaseResponse base response.
type BaseResponse struct {
	ErrCode uint        `json:"errCode"`
	ErrMsg  string      `json:"errMsg"`
	ErrDlt  string      `json:"errDlt"`
	Data    interface{} `json:"data"`
}

// UserTokenResponse user token response.
type UserTokenResponse struct {
	BaseResponse
	Data TokenData `json:"data"`
}

// TokenData token data.
type TokenData struct {
	Token             string `json:"token"`
	ExpireTimeSeconds uint   `json:"expireTimeSeconds"`
}

// BotQueryResponse bot query response.
type BotQueryResponse struct {
	Code uint         `json:"code"`
	Msg  string       `json:"msg"`
	Data BotQueryData `json:"data"`
}

// BotQueryData bot query data.
type BotQueryData struct {
	Context  string `json:"context"`
	Question string `json:"question"`
	Text     string `json:"text"`
}
