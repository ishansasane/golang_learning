package myutil

// packages are titely coupled with dirs

import "fmt"

// if package name starts with lower case it declared as a private

func PrintMessage(message string) {
	fmt.Println(message)
}
