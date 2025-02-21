package main

import "fmt"

func simpleFunctionchan() {
	fmt.Println("This is a simple function")
}

// func add(a, b int) int {
// 	return a + b
// }

// function with named return values
func namedReturnValues(a, b int) (result int) {
	result = a + b
	return
}

func main() {
	fmt.Println("Implementation fo Functions in GO")

	simpleFunctionchan()
	fmt.Println("Addition is :", namedReturnValues(2, 2))
}
