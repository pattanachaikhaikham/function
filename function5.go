package main

import "fmt"

func writeLine(a ...interface{}) {
	for _, num1 := range a {
		fmt.Println(num1)
	}
}

func main() {
	writeLine(1, 3.14, "Hello", true)
}
