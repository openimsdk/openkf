// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package api

import (
	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/common/response"
	"github.com/OpenIMSDK/OpenKF/server/internal/param"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/gin-gonic/gin"
)

// SendCode
// @Tags mail
// @Summary SendCode
// @Description Use email to send verification code
// @Produce application/json
// @Param data body param.SendToParams true "Email address"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/email/code [post]
func SendCode(c *gin.Context) {
	var params param.SendToParams
	err := c.ShouldBindJSON(&params)
	if err != nil {
		response.FailWithCode(common.INVALID_PARAMS, c)
		return
	}

	svc := service.NewServiceWithGin(c)
	err = svc.SendCode(params.Email)
	if err != nil {
		log.Debug("SendCode error: ", err)
		response.FailWithCode(common.ERROR, c)
		return
	}

	response.Success(c)
}
