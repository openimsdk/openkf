package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// define rate limit struct
type frequencyControlByTokenBucket struct {
	refreshRate float64    // 令牌的刷新速率
	capacity    int64      // 桶的容量
	tokens      float64    // 当前令牌数量
	lastToken   time.Time  // 上一次放令牌的时间
	mtx         sync.Mutex // 互斥锁
}

// allow frequency
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
	} else {
		return false
	}
}

// registried a middle ware
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
