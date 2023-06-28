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

	urltrie "github.com/OpenIMSDK/OpenKF/server/internal/middleware/hooks/url_trie"
	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/gin-gonic/gin"
)

func init() {
	urltrie.RegisterHook(MailHook{})
	fmt.Println("RegisterHook", "Register Hook[MailHook] success...")
}

type MailHook struct {
	urltrie.Hook
}

func (h MailHook) Pattern() string {
	return "/api/v1/register/email/code"
}

func (h MailHook) BeforeRun(c *gin.Context) {
	log.Debugf("GlobalHook", "path: %v", c.Request.URL.Path)
	c.Next()
}

func (h MailHook) AfterRun(c *gin.Context) {
	c.Next()
}
