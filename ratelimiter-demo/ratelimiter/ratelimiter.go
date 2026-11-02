// Package ratelimiter provides request rate limiting using the token
// bucket algorithm.
package ratelimiter

import (
	"fmt"
	"math"
	"sync"
	"time"
)

// Limiter is implemented by any rate-limiting strategy. Callers depend on
// this interface, not on TokenBucket directly — swapping in a different
// algorithm (sliding window, leaky bucket) later requires no change to
// calling code, only a new type that also implements Allow() bool.
type Limiter interface {
	Allow() bool
}

// TokenBucket implements Limiter. Tokens refill continuously at refillRate
// per second, up to capacity. Each call to Allow consumes one token if
// one is available.
type TokenBucket struct {
	mu         sync.Mutex // guards every field below; TokenBucket is safe for concurrent use
	capacity   float64
	tokens     float64
	refillRate float64 // tokens added per second
	lastRefill time.Time
}

// NewTokenBucket constructs a TokenBucket that starts completely full.
// capacity and refillRate must both be positive.
func NewTokenBucket(capacity, refillRate float64) (*TokenBucket, error) {
	if capacity <= 0 || refillRate <= 0 {
		return nil, fmt.Errorf("ratelimiter: capacity and refillRate must be positive, got capacity=%v refillRate=%v", capacity, refillRate)
	}
	return &TokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refillRate: refillRate,
		lastRefill: time.Now(),
	}, nil
}

// Allow reports whether a request may proceed right now. If it returns
// true, one token has been consumed. Safe to call from multiple goroutines.
func (t *TokenBucket) Allow() bool {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.refill()

	if t.tokens < 1 {
		return false
	}
	t.tokens--
	return true
}

// refill adds tokens based on time elapsed since the last refill.
// Callers must hold t.mu before calling this.
func (t *TokenBucket) refill() {
	now := time.Now()
	elapsed := now.Sub(t.lastRefill).Seconds()
	t.tokens = math.Min(t.capacity, t.tokens+elapsed*t.refillRate)
	t.lastRefill = now
}
