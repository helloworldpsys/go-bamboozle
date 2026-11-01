package main
import (
	"fmt"
	"math"
)

func main(){
// var declartion
var a, b int = 4, 82
var result int 

result = a + b
fmt.Println("Addition:", result)

result = a - b	
fmt.Println("Subtraction:", result)

result= a * b
fmt.Println("Multiplication:", result)

result = a / b
	fmt.Println("Division:", result)

result = a %b
fmt.Println("Remainder:", result)

const p float64 = 22.0/7.0
fmt.Println(p)

// int8 ki range -128 se 127 hoti hai.
// uint8 ki range 0 se 255 hoti hai.
// uint64 ki maximum value 18,446,744,073,709,551,615 hoti hai.

//overflow with unsigned integers
var uMaxInt uint64 = 18446744073709551615   //max value 
fmt.Println(uMaxInt)

uMaxInt = uMaxInt + 1
fmt.Println(uMaxInt)

//underflow with floating point numbers
var smallfloat float64 = 1.0e-323 
fmt.Println(smallfloat)
fmt.Println(smallfloat / math.MaxFloat64)

}