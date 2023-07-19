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

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
	"github.com/gin-gonic/gin"
)

var _ urltrie.Hook = (*CROS)(nil)

func init() {
	urltrie.RegisterHook(&CROS{})
	fmt.Println("RegisterHook", "Register Hook[CROS] success...")
}

type CROS struct {
	urltrie.Hook
}

// EDIT THIS TO YOUR OWN HOOK PATTERN
func (h *CROS) Patterns() []string {
	return []string{
		"/*",
	}
}

// EDIT THIS TO YOUR OWN HOOK PRIORITY
func (h *CROS) Priority() int64 {
	return 0
}

// EDIT THIS TO YOUR OWN HOOK BEFORE RUN, DO NOT NEED USE Next() FUNCTION
func (h *CROS) BeforeRun(c *gin.Context) {
	// c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Origin", c.GetHeader("origin"))
	c.Writer.Header().Set("Access-Control-Max-Age", "70000")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length,token")
	c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Headers,token")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
}

// EDIT THIS TO YOUR OWN HOOK AFTER RUN, DO NOT NEED USE Next() FUNCTION
func (h *CROS) AfterRun(c *gin.Context) {
}
