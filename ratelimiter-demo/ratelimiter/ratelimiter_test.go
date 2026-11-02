package ratelimiter

import "testing"

func TestTokenBucket_Allow(t *testing.T) {
	tb, err := NewTokenBucket(2, 1) // capacity 2, refills 1 token/sec
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	tests := []struct {
		name string
		want bool
	}{
		{"first request consumes a token", true},
		{"second request consumes the last token", true},
		{"third request has no tokens left", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tb.Allow(); got != tt.want {
				t.Errorf("Allow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewTokenBucket_RejectsInvalidInput(t *testing.T) {
	if _, err := NewTokenBucket(-1, 1); err == nil {
		t.Error("expected an error for negative capacity, got nil")
	}
}
