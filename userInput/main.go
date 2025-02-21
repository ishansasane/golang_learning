package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey ! whats Your Name :")
	// var name string

	// fmt.Scan(&name)

	// fmt.Printf("Input is : %s\n", name)

	// reading from buffer reader
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Printf("Input is : %s\n", name)
}
