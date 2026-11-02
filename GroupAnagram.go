package main

import (
"fmt"
"sort"
"time"
)

func sortString(s string) string{
	b := []byte(s)
	sort.Slice(b, func(i,j int)bool{
		return b[i]<b[j]
	})
	return string(b)
}

func groupAnagrams(strs []string) [][]string{
	groups := make(map[string][]string)

	for _, s := range strs{
		key := sortString(s)
		groups[key] = append(groups[key],s)
	}

	//var result [][]string
	result := make([][]string, 0, len(groups))

	for _,v := range groups{
		result = append(result,v)
	}
	return result
}

func main(){

	strs := []string{"cat","act","tac","dog","god","bat","tab","red"}
	start := time.Now()
	res := groupAnagrams(strs)
	elapsed := time.Since(start)
	fmt.Println("Anagram result:", res)
	fmt.Println("Time taken:", elapsed)
}