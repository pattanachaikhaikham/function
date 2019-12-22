package main

import "fmt"

func suntract(number int) {
	number = number - 1
}

func main() {
	num1 := 10
	suntract(num1)
	fmt.Print(num1)
}
