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

package service

import (
	"time"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/client"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

// SendCode send code to email
func (svc *Service) SendCode(email string) (err error) {
	// check the code is exist
	cmd := svc.Cache.Get(svc.Ctx, "code:"+email)
	var code string
	if cmd.Err() != nil {
		code = utils.GenerateCode()
		// save code in 60s
		status := svc.Cache.Set(svc.Ctx, "code:"+email, code, time.Second*60)
		if status.Err() != nil {
			return status.Err()
		}
	}

	// generate code
	err = client.SendEmail(email, "OpenKF Admin Register", "Your verification code is "+code)

	return err
}

// Test test
func (svc *Service) Test() {
	log.Info("test...")
}
