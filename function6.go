package main

import "fmt"

func main() {
	add := func(num1, num2 int) int {
		return num1 + num2
	}
	num1 := add(10, 20)
	fmt.Println(num1)
}
