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

package auth

import (
	"fmt"

	"github.com/openimsdk/openkf/server/pkg/openim/client"
	"github.com/openimsdk/openkf/server/pkg/openim/param/request"
	"github.com/openimsdk/openkf/server/pkg/openim/param/response"
)

const (
	// pathUserToken get user token path.
	pathUserToken = "/auth/user_token"
)

// GetUserToken get user token from openim server.
func GetUserToken(param *request.UserTokenParams, operationID, host string) (*response.UserTokenResponse, error) {
	// host: http://ip:port
	url := fmt.Sprintf("%s%s", host, pathUserToken)

	r := &response.UserTokenResponse{}
	client := client.NewClient(url)
	resp, err := client.POST(operationID, "", param)
	if err != nil {
		return r, err
	}

	code, ok := resp["errCode"].(float64)
	if !ok {
		return r, fmt.Errorf("code is not float64")
	}
	r.ErrCode = uint(code)

	r.ErrMsg, ok = resp["errMsg"].(string)
	if !ok {
		return r, fmt.Errorf("msg is not string")
	}

	r.ErrDlt, ok = resp["errDlt"].(string)
	if !ok {
		return r, fmt.Errorf("msg is not string")
	}

	if data, ok := resp["data"]; ok {
		data, ok := data.(map[string]interface{})
		if !ok {
			return r, fmt.Errorf("data is not map[string]interface{}")
		}

		r.Data.Token, ok = data["token"].(string)
		if !ok {
			return r, fmt.Errorf("token is not string")
		}

		expireTimeSeconds, ok := data["expireTimeSeconds"].(float64)
		if !ok {
			return r, fmt.Errorf("expireTimeSeconds is not float64")
		}

		r.Data.ExpireTimeSeconds = uint(expireTimeSeconds)
	}

	return r, nil
}
