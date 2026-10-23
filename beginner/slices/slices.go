package main 

import "fmt"
//slice are dynamic array 
// useful methods
func main(){
	//uninitalized slice is nill
var nums[]int
fmt.Println(nums == nil)

a := [5]int{10, 20, 30, 40, 50}

s := a[1:4]

a[0] = 999

fmt.Println(a[0])

fmt.Println(s)



}