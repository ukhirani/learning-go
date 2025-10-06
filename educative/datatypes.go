// package main

import (
	fm "fmt"
)

func main() {
	var hello int = 6 // this is how you declare a variable -> var, variable name,data type

	//elementary = primitive
	//structured = composite -> default value before initializng = nil

	fm.Println(hello)

	//functions can also have a datatype and their datatype is the datatype they return

	var x, y int
	x, y = dummy1(1, 2)

	fm.Println(x, y)

	type IZ int

	type (
		IX  int
		FZ  float32
		STR string
	)

	var xx IX
	xx = 5
	print(xx)

}

func dummy(a int, b string) int {
	return 0
}

func dummy1(a int, b int) (x int, y int) {
	x = a + b
	y = a * b
	return x, y
}
