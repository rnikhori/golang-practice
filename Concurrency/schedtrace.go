package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func busy(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	for time.Since(start) < 3*time.Second {
		// CPU ghuma rahe hain, koi block nahi
	}
}

func main() {
	runtime.GOMAXPROCS(4)

	fmt.Println("phase 1: main 1 second so raha hai, koi kaam nahi")
	time.Sleep(1 * time.Second)

	fmt.Println("phase 2: 8 busy goroutines, sirf 4 P")
	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go busy(i, &wg)
	}
	wg.Wait()

	fmt.Println("phase 3: sab khatam, 1 second phir se khaali")
	time.Sleep(1 * time.Second)
}
