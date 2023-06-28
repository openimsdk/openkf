// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package response

import (
	"net/http"

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewResponse(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Success(c *gin.Context) {
	NewResponse(common.SUCCESS, common.GetMsg(common.SUCCESS), nil, c)
}

func SuccessWithData(data interface{}, c *gin.Context) {
	NewResponse(common.SUCCESS, common.GetMsg(common.SUCCESS), data, c)
}

func SuccessWithCode(code int, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), nil, c)
}

func SuccessWithAll(code int, data interface{}, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), data, c)
}

func Fail(c *gin.Context) {
	NewResponse(common.ERROR, common.GetMsg(common.ERROR), nil, c)
}

func FailWithData(data interface{}, c *gin.Context) {
	NewResponse(common.ERROR, common.GetMsg(common.ERROR), data, c)
}

func FailWithCode(code int, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), nil, c)
}

func FailWithAll(code int, data interface{}, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), data, c)
}
