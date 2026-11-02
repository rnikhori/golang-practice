package main
import "fmt"

func main(){
	count:=0
	for {
		if count > 10 {
			break
		}
		fmt.Println(count)
		count++
		if count%2 == 0 {
			continue
		} 
		
		
	}
}