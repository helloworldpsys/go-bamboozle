package main

import "fmt"

func calculate(a, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := calculate(10, 3)

    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)

	quotient, _ = calculate(10,3)
	fmt.Println(quotient)

	_ = quotient

}