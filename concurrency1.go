package main

import (
	"fmt"
	"sync"
	"time"
) 

func fetch(id int ,wg *sync.WaitGroup) {
	wg.Done()
	fmt.Println("fetching <id>", id)
	time.Sleep(1 * time.Second)
}
func main(){
	start := time.Now()
	var wg sync.WaitGroup
	for i:=0; i<5; i++ {
		wg.Add(i)
		go fetch(i,wg &sync.WaitGroup{})
	}
	wg.Wait()
	fmt.Println(time.Since(start))
}

