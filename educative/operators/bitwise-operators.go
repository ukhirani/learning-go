package main

import "fmt"

func main() {
	x := 3
	y := 5
	z := (x ^ y) & (x << 2)
	fmt.Println(z)
}
