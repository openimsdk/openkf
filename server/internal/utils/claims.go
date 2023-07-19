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

package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GetUserID get user uuid from context.
func GetUserUUID(c *gin.Context) (string, error) {
	if claims, ok := c.Get("claims"); ok {
		if c := claims.(*JwtClaims); c != nil {
			return c.UserUUID, nil
		}
	}

	return "", errors.Errorf("get user uuid failed")
}

// GetCommunityUUID get community uuid from context.
func GetCommunityUUID(c *gin.Context) (string, error) {
	if claims, ok := c.Get("claims"); ok {
		if c := claims.(*JwtClaims); c != nil {
			return c.CommunityUUID, nil
		}
	}

	return "", errors.Errorf("get community uuid failed")
}
