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

	"github.com/gin-gonic/gin"

	"github.com/openimsdk/openkf/server/pkg/log"
)

// Mapping for store url pattern and hook.
var hookTrie *Trie

func init() {
	hookTrie = NewTrie()
}

// RegisterHook register url & hook to trie.
func RegisterHook(hook Hook) {
	hookTrie.InsertBatch(hook.Patterns(), hook)
}

// RunHook enable hook for interceptor.
func RunHook() gin.HandlerFunc {
	return func(c *gin.Context) {
		raw := c.Request.URL.Path

		// Get path from url
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

		// Run all before hooks
		for _, hook := range hooks {
			hook.BeforeRun(c)
		}

		// Run controllers
		c.Next()

		// Run all after hooks
		for _, hook := range hooks {
			hook.AfterRun(c)
		}
	}
}
