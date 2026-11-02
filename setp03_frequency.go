package main

import "fmt"


func main(){

	nums := []int{3,1,3,2,1,3,4}
	freq := make(map[int]int)

	for i,n := range nums {
		freq[n]++
		fmt.Printf("step %d: saw %d -> %v\n",i+1, n, freq)
	}

	fmt.Println("\nFinal:",freq)
}

