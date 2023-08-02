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
	"mime/multipart"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/internal/conn/client"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
)

// CommonService community service.
type CommonService struct {
	Service
}

// NewCommonService return new service with gin context.
func NewCommonService(c *gin.Context) *CommonService {
	return &CommonService{
		Service: Service{
			ctx: c,
		},
	}
}

// UploadFile upload file.
func (svc *CommonService) UploadFile(file *multipart.FileHeader) (string, error) {
	fileHandler, err := file.Open()
	if err != nil {
		return "", err
	}
	defer fileHandler.Close()

	filename := file.Filename
	objectName := utils.GenerateObjectName(filename)
	fileSize := file.Size

	err = client.PutObject(objectName, fileHandler, fileSize)
	if err != nil {
		return "", err
	}

	url := client.GetObejctURL(objectName)

	return url, nil
}
