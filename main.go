// go mod init for external files for go/src

// main package important main function
package main

// file name dosent matter in golang just main function matters for execution

//special package
import (
	"fmt"
)

// main function is a executable function and it must be a part of main package

func main() {
	fmt.Println("Hello World")

	var name string = "Ishan"
	var version = 1.1
	fmt.Println(version)
	fmt.Println(name)

	ishan := "Sasane"
	fmt.Println(ishan)
	// for getting type of var
	fmt.Printf("Type of name is %T\n", name)

}

// for running it - > go run file_name
