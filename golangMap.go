package main

import (
	"fmt"
)

func main(){
	//A map stores a key -> value pair for example : name -> price
	prices := make(map[string]int)

	// insert 
	prices["apple"]= 100
	prices["banana"] = 40
	prices["mango"] = 140
	//fmt.Println("After insert:", prices)

	//Read

	//fmt.Println("Price of an apple is:",prices["apple"])

	//update 
	prices["apple"]= 120
	fmt.Println("After update:", prices)

	//length of keys or number of keys in the map
	fmt.Println("Number of fruits:",len(prices))

	delete(prices,"banana")

	fmt.Println("After deleting banana key ,now prices:",prices)

	prices["grapes"]= 90
	fmt.Println("Now the map looks like this",prices)
}