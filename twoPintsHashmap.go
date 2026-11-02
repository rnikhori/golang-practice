package main
import "fmt"

func targetSum(nums []int,target int) []int{
	seen := map[int]int{}
	for i, num := range nums{
		complement := target - num 
		if idx,ok := seen[complement]; ok{
			return []int{idx,i}
		}
		seen[num]= i
	}
	return []int{}
}
func main(){
	nums := []int{2,7,11,15}
	target := 9
	result := targetSum(nums, target)
	fmt.Println(result)
}