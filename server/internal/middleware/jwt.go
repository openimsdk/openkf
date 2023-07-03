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

package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
)

// EnableAuth enable auth middleware.
func EnableAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" || strings.Fields(token)[0] != "Bearer" {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		_, err := utils.ParseJwtToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)

			return
		}

		// todo: add info to claims
		c.Next()
	}
}
