package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)
type Event struct {
	mu sync.Mutex
	fired bool
	cbs []func()
}

func (e *Event) Register(cb func()){
	e.mu.Lock()
	if !e.fired {
		e.cbs = append(e.cbs, cb)
		e.mu.Unlock()
		return;
	}
	e.mu.Unlock()
	cb()
}

func (e *Event) Fire() {
    e.mu.Lock()
    if e.fired { // doosri baar: kuch nahi
        e.mu.Unlock()
        return
    }
    e.fired = true
    toRun := e.cbs // list nikaal li
    e.cbs = nil
    e.mu.Unlock()

    for _, cb := range toRun { // lock ke bahar chalao
        cb()
    }
}

func main() {
    var e Event
    var count int64
    var wg sync.WaitGroup

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            e.Register(func() { atomic.AddInt64(&count, 1) })
        }()
    }
    wg.Add(1)
    go func() { defer wg.Done(); e.Fire() }()

    wg.Wait()
    fmt.Println("count =", count) // 100 aana chahiye
}