package main

import (
	"fmt"

)

func main(){
	nums := []int{6, 7, 8}
	sum := 0

	for _, num := range nums {
		sum = sum +num
	}
	fmt.Println(sum)

	m := map[string]string{
		"fisrt name": "kaim",
		"last name": "kaimx"}
		for key , value := range m {
			fmt.Println(key, value)
			}
			fo
} 