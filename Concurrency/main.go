package main

import (
	"fmt"
	"sync"
	"time"
)

type result struct{
	Value string
	Err error
}

func worker(url string , wg *sync.WaitGroup,resultchan chan result) {
	defer wg.Done()
	time.Sleep(time.Millisecond * 50)
	fmt.Printf("image processed : %s\n",url)
	resultchan <- result{
		Value: url,
		Err: nil,
	}

}

func main(){
//	fmt.Print("welcome to go concurrency")
		var wg sync.WaitGroup
		resultchan := make(chan result,5)
		wg.Add(5)
	 startTime := time.Now()
	 go worker("image_1.png",&wg ,resultchan )
	 go worker("image_2.png", &wg ,resultchan )
	 go worker("image_3.png",&wg ,resultchan )
	 go worker("image_4.png", &wg ,resultchan )
	 go worker("image_5.png", &wg ,resultchan )

	 wg.Wait()
	 close(resultchan)
	 for result:= range resultchan {
		fmt.Printf("recieved image. %v \n",result)
		if result.Err != nil{
			// reprocess the queue -DLQ
		}
	 }
	fmt.Printf("it took %s ms. \n",time.Since(startTime))
}