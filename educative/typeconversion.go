// package main

import "fmt"

//go does not allow for implicit type casting
//it only allows for explicit type casting so the code can figure out

func main() {
	var number float32 = 5.2 // Declared a floating point variable
	fmt.Println(number)      // Printing the value of variable
	fmt.Println(int(number)) // Printing the type-casted result
}
