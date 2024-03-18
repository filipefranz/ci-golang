package main

import "fmt"

func main() {
	fmt.Println(Soma(10, 20))
}

func Soma(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mult(a, b int) int {
	return a * b
}

func Div(a, b int) int {
	return a / b
}
