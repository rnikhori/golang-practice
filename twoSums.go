package main
import "fmt"

func twoSums(nums []int, target int) []int {

	//[2,7,11,15] target = 9 
	
	n := len(nums)
	for i:=0; i< n; i++ {
		for j := i+1; j<n; j++ {
			if nums[i]+ nums[j] == target {
				return []int{i,j}
			}
		}
	}
	return []int{}
}

func main(){
	nums := []int{2,7,11,15}
	target := 9
	result := twoSums(nums, target)
	fmt.Println(result)	
}