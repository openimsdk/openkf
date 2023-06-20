// Copyright © 2023 OpenIMSDK open source community. All rights reserved.
// Licensed under the MIT License (the "License");
// you may not use this file except in compliance with the License.

package middleware

import (
	"net/http"
	"strings"

	"github.com/OpenIMSDK/OpenKF/server/pkg/utils"
	"github.com/gin-gonic/gin"
)

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
