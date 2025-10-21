package main

import "fmt"

func main() {
	// A value that can be changed by a program during execution is called a variable

	// var identifier type
	// As discussed earlier in this chapter, type is written after the identifier of the variable, contrary to most older programming languages.

	// When a variable is declared, memory in Go is initialized, which means it contains the default zero or null value depending upon its type automatically.

	var num int      // Declaring  an integer variable
	fmt.Println(num) // Printing its value

	var decision bool     // Declaring a boolean variable
	fmt.Println(decision) // Printing its value

	// For example, 0 for int, 0.0 for float, false for bool, empty string ("") for string, nil for pointer, zero-ed struct, and so on.

	// go run variables.go
	// 0
	// false

	fmt.Println("wohoo")

	// Remark: The naming of identifiers for variables follows the camelCasing rules (start with a small letter, and every new part of the word starts with a capital letter). But if the variable has to be exported, it must start with a capital letter, as discussed earlier in this chapter.
}
