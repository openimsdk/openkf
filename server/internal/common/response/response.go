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

import (
	"net/http"

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/gin-gonic/gin"
)

// Response is a common struct for response
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// NewResponse returns a new Response
func NewResponse(code int, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, &Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

// Success returns a success response
func Success(c *gin.Context) {
	NewResponse(common.SUCCESS, common.GetMsg(common.SUCCESS), nil, c)
}

// SuccessWithData returns a success response with data
func SuccessWithData(data interface{}, c *gin.Context) {
	NewResponse(common.SUCCESS, common.GetMsg(common.SUCCESS), data, c)
}

// SuccessWithCode returns a success response with code
func SuccessWithCode(code int, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), nil, c)
}

// SuccessWithAll returns a success response with code and data
func SuccessWithAll(code int, data interface{}, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), data, c)
}

// Fail returns a fail response
func Fail(c *gin.Context) {
	NewResponse(common.ERROR, common.GetMsg(common.ERROR), nil, c)
}

// FailWithData returns a fail response with data
func FailWithData(data interface{}, c *gin.Context) {
	NewResponse(common.ERROR, common.GetMsg(common.ERROR), data, c)
}

// FailWithCode returns a fail response with code
func FailWithCode(code int, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), nil, c)
}

// FailWithAll returns a fail response with code and data
func FailWithAll(code int, data interface{}, c *gin.Context) {
	NewResponse(code, common.GetMsg(code), data, c)
}
