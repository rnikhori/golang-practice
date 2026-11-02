package main 
import "fmt"

func duplicateCheck(nums []int) bool{
	dmap := map[int]int{}
	 {
		if _,ok := dmap[num];ok{
			return true
		}
		dmap[num] = i

	}
	return false
}

func main(){
	nums := []int{1,2,3,1}
	result := duplicateCheck(nums)
	fmt.Println(result)
}