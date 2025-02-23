package main

import (
	"fmt"
)

func modify(num *int) {
	*num = 100
}

func main() {
	// var num int
	// num = 4

	// var ptr *int
	// ptr = &num

	num := 4
	ptr := &num
	modify(&num)

	fmt.Println(ptr, *ptr)

}
