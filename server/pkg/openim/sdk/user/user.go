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

package user

import (
	"fmt"

	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/client"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/request"
	"github.com/OpenIMSDK/OpenKF/server/pkg/openim/param/response"
)

const (
	// PATH_USER_REGISTER register user path.
	PATH_USER_REGISTER = "/auth/user_register"
)

// RegisterUser register user.
func RegisterUser(param *request.RegisterUserParams, host string) (*response.BaseResponse, error) {
	// host: http://ip:port
	url := fmt.Sprintf("%s%s", host, PATH_USER_REGISTER)

	r := &response.BaseResponse{}
	client := client.NewClient(url)
	resp, err := client.POST(PATH_USER_REGISTER, param)
	if err != nil {
		return r, err
	}

	r.ErrCode = uint(resp["errCode"].(float64))
	r.ErrMsg = resp["errMsg"].(string)
	r.ErrDlt = resp["errDlt"].(string)

	return r, nil
}
