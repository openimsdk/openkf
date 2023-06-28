// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package service

import (
	"time"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/client"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
)

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

func (svc *Service) Test() {
	log.Info("test...")
}
