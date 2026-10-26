package main

import (
	"fmt"
	"slices"
)

//slices are dynamic arrayss no need to give it length it can expand automatically
//most used construct in go
//we can useful methods in go


func main() {
//unintialized slice is nil
	// var nums []int
	// fmt.Println(nums == nil)
	// fmt.Println(len(nums))

	// intialized 
	//2 is a intial size
	// 5 is initial capcity
	var nums = make ([]int, 2, 5)// output[0,0] default value
	fmt.Println(nums)
	fmt.Println(cap(nums)) // to check the capacity--> maximum number can fit the output will be 2
	nums = append(nums, 1)
	fmt.Println(nums)// output [0,0,1]
	nums = append(nums, 2)//output [0,0,1,2] abhi tak capacity will be 5 but now when i add 3 and 4 capacity will double 5 se 10 




	a := [6]int{0, 1, 2, 3, 4, 5}
	var slice = a[1:6]
	// a[:]makes a slice

	slice = append(slice, a[2], 10)
	slice = append(slice, a[:]...)

	//a[:]... unpacks the array into indidual elements

	var slice_copy = make([]int, len(a))
	copy(slice_copy, a[:])

	//or

	// var slice_copy [6]int
	// copy(slice_copy[:],a[:])
 
	fmt.Println(slice)
	fmt.Println(slice_copy)

	var slice_equal = slices.Equal(slice, slice_copy)

	fmt.Println("slice and slice copy equal:", slice_equal)

	var twoD = make([][]int, 4)

	for i := 0; i < 4; i++ {
		var innerLength = i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

	slice = slice[2:4]

	fmt.Println("capacity of slice is", cap(slice))
	// slice is a portion of array amd cap tells the capacity of the array

	fmt.Println("length of slice is", len(slice))
	//length is how big the array is!

	//slice operator
	// var nums = []int{1,2,3}
	// fmt.Println(nums[0:2])


//from slices pacakage
	var nums1 = []int{1,2,3}
	var nums2 = []int{1,2,3}
	fmt.Println(slices.Equal(nums1, nums2))

	//we can make 2d slices
	var nums3 = [][]int{{1,2,3},{4,5,6}}
	fmt.Println(nums3)

}