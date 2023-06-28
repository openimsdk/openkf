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

package urltrie

import (
	"net/url"

	"github.com/OpenIMSDK/OpenKF/server/pkg/log"
	"github.com/gin-gonic/gin"
)

// Mapping for store url pattern and hook
var hookTrie *Trie

func init() {
	hookTrie = NewTrie()
}

// Register url & hook to trie
func RegisterHook(hook Hook) {
	hookTrie.Insert(hook.Pattern(), hook)
}

// Enable hook for interceptor
func RunHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.Request.URL.Path

		// get path from url
		p, err := url.Parse(raw)
		if err != nil {
			log.Errorf("Hook", "parse url error: %v", err)
		}

		path := p.Path
		hooks, ok := hookTrie.Match(path)
		if !ok {
			c.Next()
			return
		}

		// run hooks
		for _, hook := range hooks {
			hook.BeforeRun(c)
			c.Next()
			hook.AfterRun(c)
		}
	}
}
