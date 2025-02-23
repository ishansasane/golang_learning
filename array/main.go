package main

import "fmt"

func main() {
	// we are learning array in golang
	var name [5]string
	name[0] = "John"
	name[1] = "Doe"
	name[2] = "John"
	name[3] = "Doe"
	name[4] = "John"
	fmt.Println(name)
	fmt.Println(len(name))

	var numbers = [5]int{1, 2, 3, 4}
	fmt.Println(len(numbers))

	// create one for loop
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}

}
