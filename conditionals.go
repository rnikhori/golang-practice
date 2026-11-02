package main
import "fmt"

func main(){
	score := 75
	if score >= 90 || score >= 75 {
		fmt.Println("A")
	}else if score >= 75 {
		fmt.Println("B")
	}else if score >= 50 {
		fmt.Println("C")
	}else{
		fmt.Println("F")
	}
}