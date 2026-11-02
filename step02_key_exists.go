package main

import(
	"fmt"
)

func hasKey(m map[string]int, k string)bool{
	if _,ok := m[k]; ok{
		fmt.Println("Got the key in map\n",k)
		return true
	} else{
		fmt.Println("Not found the  key in map\n",k)
		return false
	}
}

func main(){
	stock := map[string]int{
		"pen":10,
		"pencil":0,
	}

	fmt.Println("pencil:", stock["pencil"])
	fmt.Println("pen:", stock["pen"])

	fmt.Println()

	for _,item := range []string{"pen","pencil","eraser","sharpner"} {
		// if qty ,ok := stock[item]; ok {
		// 	fmt.Printf("%-7s found,qty = %d\n",item,qty)
		// } else{
		// 	fmt.Printf("%-7s not found\n",item)
		// }
		hasKey(stock,item)
	}
}