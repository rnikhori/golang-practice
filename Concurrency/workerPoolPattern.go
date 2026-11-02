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

func worker(jobsChan chan string , wg *sync.WaitGroup,resultchan chan result) {
	defer wg.Done()
	for job:= range jobsChan {
			time.Sleep(time.Millisecond * 50)
			//fmt.Printf("image processed : %s\n",job)
			resultchan <- result{
		Value: job,
		Err: nil,
	}
	 }

	 fmt.Println("worker shutting down")


	

}

func main(){
//	fmt.Print("welcome to go concurrency")
jobs := []string{
	"image_1.png", "image_2.png", "image_3.png", "image_4.png", "image_5.png",
	"image_6.png", "image_7.png", "image_8.png", "image_9.png", "image_10.png",
	"image_11.png", "image_12.png", "image_13.png", "image_14.png", "image_15.png",
	"image_16.png", "image_17.png", "image_18.png", "image_19.png", "image_20.png",
}
		var wg sync.WaitGroup
		totalWorkers := 5
		resultchan := make(chan result,50)
		jobsChan := make(chan string, len(jobs))
	
	 startTime := time.Now()
	 for i:=1; i<=totalWorkers; i++{
		wg.Add(1)
		go worker(jobsChan, &wg, resultchan)
	 }

	 go func (){
		 wg.Wait()
	 close(resultchan)
	 }(	)

	 //send the jobs
	 for i:=0;i<len(jobs);i++{
		jobsChan <- jobs[i]
	 }

	 close(jobsChan)
	
	 for result:= range resultchan {
		fmt.Printf("Job completed %v \n",result)
	 }
	fmt.Printf("it took %s ms. \n",time.Since(startTime))
}