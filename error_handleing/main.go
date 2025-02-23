package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error: Division by zero")
	}
	return a / b, nil

}

func main() {
	// error handleing
	fmt.Println("Error Handleing")
	result, err := divide(10, 0)
	fmt.Println("result is :", result, err)

}
