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
}