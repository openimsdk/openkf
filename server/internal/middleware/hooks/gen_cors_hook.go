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

	"github.com/gin-gonic/gin"

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
)

var _ urltrie.Hook = (*CORS)(nil)

func init() {
	urltrie.RegisterHook(&CORS{})
	fmt.Println("RegisterHook", "Register Hook[CORS] success...")
}

// CORS implement urltrie.Hook.
type CORS struct {
	urltrie.Hook
}

// Patterns EDIT THIS TO YOUR OWN HOOK PATTERN.
func (h *CORS) Patterns() []string {
	return []string{
		"/*",
	}
}

// Priority EDIT THIS TO YOUR OWN HOOK PRIORITY.
func (h *CORS) Priority() int64 {
	return 0
}

// BeforeRun EDIT THIS TO YOUR OWN HOOK BEFORE RUN, DO NOT NEED USE Next() FUNCTION.
func (h *CORS) BeforeRun(c *gin.Context) {
	method := c.Request.Method
	origin := c.GetHeader("Origin")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
	c.Header("Access-Control-Allow-Credentials", "true")

	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

// AfterRun EDIT THIS TO YOUR OWN HOOK AFTER RUN, DO NOT NEED USE Next() FUNCTION.
func (h *CORS) AfterRun(c *gin.Context) {
}
