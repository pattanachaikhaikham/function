package main

import "fmt"

func sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total = total + n
	}
	return total
}
func main() {
	num1 := sum(1, 3, 5, 7, 9)
	fmt.Println(num1)

	num2 := sum()
	fmt.Println(num2)
}
