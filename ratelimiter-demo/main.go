package main

import (
	"fmt"
	"log"

	"ratelimiter-demo/ratelimiter"
)

func main() {
	limiter, err := ratelimiter.NewTokenBucket(3, 1) // 3 tokens capacity, refills 1/sec
	if err != nil {
		log.Fatalf("failed to create limiter: %v", err)
	}

	for i := 1; i <= 5; i++ {
		if limiter.Allow() {
			fmt.Printf("request %d: allowed\n", i)
		} else {
			fmt.Printf("request %d: rejected (rate limited)\n", i)
		}
	}
}
