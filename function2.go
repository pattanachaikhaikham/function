package main

import "fmt"

func num(num1 int, num2 int) int {
	ans := num1 - num2
	return ans
}

func main() {
	num1 := num(5, 5)
	fmt.Println(num1)
}
