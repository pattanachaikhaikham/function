package main

import "fmt"

func say(greet string) func(string) string {
	return func(name string) string {
		return greet + name
	}
}

func main() {
	num1 := say("Hello")
	fmt.Println(num1("Goku"))
	fmt.Println(num1("Gohan"))
}
