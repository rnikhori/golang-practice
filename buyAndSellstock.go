// given an array lets say stocks = [7,1,3,5,6,2]
package main;

func maxProfit(prices [] int) int {

	n := len(prices)
	minSoFar := prices[0]
	res := 0

	for i:=1; i<n; i++{
		minSoFar = min(minSoFar,prices[i])
		res = max(res, prices[i]-minSoFar)
	}
	return res
}

func main(){

}