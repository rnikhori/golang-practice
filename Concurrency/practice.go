package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	fmt.Println("goroutine: main ne mujhe banaya")
	// }()
		for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Println("goroutine", id)
		}(i)
	}

	fmt.Println("main: goroutine bana di, ab wait karunga")
	wg.Wait()
	fmt.Println("main: goroutine khatam, ab main bhi khatam")
}