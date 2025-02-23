package main

import "fmt"

func main() {
	// slices is a better version of array is is a dynamic in size
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 9, 0)
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// using make function
	names := make([]string, 5, 5)
	names = append(names, "Ishan")
	fmt.Println(cap(names))
}
