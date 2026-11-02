package main

import (
	"fmt"
)
//step 1 creation of node 
type Node struct {
	value int
	next *Node
}

//step2 linkedlist creation
type LinkedList struct{
	head *Node
}

//Linked list implementation
//define a struct in golang 
func main(){
	fmt.Println("hello world !!")
}