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
	"sync"
	"time"

	"github.com/gin-gonic/gin"

	urltrie "github.com/openimsdk/openkf/server/internal/middleware/hooks/url_trie"
)

var _ urltrie.Hook = (*RATE)(nil)

func init() {
	urltrie.RegisterHook(&RATE{})
	fmt.Println("RegisterHook", "Register Hook[RATE] success...")
}

// RATE implement urltrie.Hook.
type RATE struct {
	urltrie.Hook
}

// Patterns EDIT THIS TO YOUR OWN HOOK PATTERN.
func (h *RATE) Patterns() []string {
	return []string{
		"",
	}
}

// Priority EDIT THIS TO YOUR OWN HOOK PRIORITY.
func (h *RATE) Priority() int64 {
	return 0
}

// BeforeRun EDIT THIS TO YOUR OWN HOOK BEFORE RUN, DO NOT NEED USE Next() FUNCTION.
func (h *RATE) BeforeRun(c *gin.Context) {
}

// AfterRun EDIT THIS TO YOUR OWN HOOK AFTER RUN, DO NOT NEED USE Next() FUNCTION.
func (h *RATE) AfterRun(c *gin.Context) {
}

type frequencyControlByTokenBucket struct {
	refreshRate float64    // token refresh rate
	capacity    int64      // bucket capacity
	tokens      float64    // token count
	lastToken   time.Time  // latest time token stored
	mtx         sync.Mutex // mutex
}

// allow frequency.
func (tb *frequencyControlByTokenBucket) Allow() bool {
	tb.mtx.Lock()
	defer tb.mtx.Unlock()
	now := time.Now()
	// compute tokens which needs
	tb.tokens = tb.tokens + tb.refreshRate*now.Sub(tb.lastToken).Seconds()
	if tb.tokens > float64(tb.capacity) {
		tb.tokens = float64(tb.capacity)
	}
	// judge weather to pass through
	if tb.tokens >= 1 {
		tb.tokens--
		tb.lastToken = now

		return true
	}

	return false
}

// LimitHandler registried a middle ware to use.
func LimitHandler(maxConn int, refreshRate float64) gin.HandlerFunc {
	tb := &frequencyControlByTokenBucket{
		capacity:    int64(maxConn),
		refreshRate: refreshRate,
		tokens:      0,
		lastToken:   time.Now(),
	}

	return func(c *gin.Context) {
		if !tb.Allow() {
			c.String(503, "Too many request")
			c.Abort()

			return
		}
		c.Next()
	}
}
