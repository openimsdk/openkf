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

package pkg

import "html/template"

func checkTemplate(t string) *template.Template {
	return template.Must(template.New("").Parse(t))
}

var hookTemplate = checkTemplate(`
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
	"github.com/gin-gonic/gin"
)

var _ urltrie.Hook = (*{{.HookName}})(nil)

func init() {
	urltrie.RegisterHook(&{{.HookName}}{})
	fmt.Println("RegisterHook", "Register Hook[{{.HookName}}] success...")
}

type {{.HookName}} struct {
	urltrie.Hook
}

// Patterns EDIT THIS TO YOUR OWN HOOK PATTERN
func (h {{.HookName}}) Patterns() string {
	return "{{.UrlPattern}}"
}

// Priority EDIT THIS TO YOUR OWN HOOK PRIORITY
func (h GlobalHook) Priority() int64 {
	return {{.Prority}}
}

// BeforeRun EDIT THIS TO YOUR OWN HOOK BEFORE RUN
func (h {{.HookName}}) BeforeRun(c *gin.Context) {
	c.Next()
}

// AfterRun EDIT THIS TO YOUR OWN HOOK AFTER RUN
func (h {{.HookName}}) AfterRun(c *gin.Context) {
	c.Next()
}
`)
