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

package hooks

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
	"github.com/OpenIMSDK/OpenKF/server/internal/utils"
)

var _ urltrie.Hook = (*JWT)(nil)

func init() {
	urltrie.RegisterHook(&JWT{})
	fmt.Println("RegisterHook", "Register Hook[JWT] success...")
}

// JWT implement urltrie.Hook.
type JWT struct {
	urltrie.Hook
}

// Patterns EDIT THIS TO YOUR OWN HOOK PATTERN.
func (h *JWT) Patterns() []string {
	return []string{
		"/api/v1/user/*",
		"/api/v1/community/*",
		"/api/v1/admin/*",
		"/api/v1/platform/*",
		"/api/v1/login/exit", // temp, will remove to OpenIM Callback
	}
}

// Priority EDIT THIS TO YOUR OWN HOOK PRIORITY.
func (h *JWT) Priority() int64 {
	return 0
}

// BeforeRun EDIT THIS TO YOUR OWN HOOK BEFORE RUN, DO NOT NEED USE Next() FUNCTION.
func (h *JWT) BeforeRun(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" || strings.Fields(token)[0] != "Bearer" {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	claims, err := utils.ParseJwtToken(strings.Fields(token)[1])
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)

		return
	}

	// Set claims to context.
	c.Set("claims", claims)
}

// AfterRun EDIT THIS TO YOUR OWN HOOK AFTER RUN, DO NOT NEED USE Next() FUNCTION.
func (h *JWT) AfterRun(c *gin.Context) {
}
