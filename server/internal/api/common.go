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

package api

import (
	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/common"
	"github.com/OpenIMSDK/OpenKF/server/internal/common/response"
	requestparams "github.com/OpenIMSDK/OpenKF/server/internal/params/request"
	"github.com/OpenIMSDK/OpenKF/server/internal/service"
)

// UploadFile
// @Tags common
// @Summary UploadFile
// @Description upload a file
// @Produce application/json
// @Param file formData file true "file"
// @Success 200 {object}  response.Response{msg=string} "Success"
// @Router /api/v1/common/file/upload [post].
func UploadFile(c *gin.Context) {
	file, err := c.FormFile(requestparams.FILE_REQUEST_PARAMS)
	if err != nil {
		response.FailWithAll(common.INVALID_PARAMS, err.Error(), c)

		return
	}

	svc := service.NewCommonService(c)
	url, err := svc.UploadFile(file)
	if err != nil {
		response.FailWithAll(common.KF_INTERNAL_ERROR, err.Error(), c)

		return
	}

	response.SuccessWithData(url, c)
}
