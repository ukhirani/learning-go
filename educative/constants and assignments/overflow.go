// package main

import fm "fmt"

func main() {

	//

	const Ln2 = 0.693147180559945309417232121458176568075500134360255254120680009
	const Log2E = 1 / Ln2 // this is a precise reciprocal
	const BILLION = 1e9   // float constant
	const HARD_EIGHT = (1 << 100) >> 97

	fm.Println(Ln2, Log2E, BILLION, HARD_EIGHT)

	//here's what the terminal has to say

	//go run overflow.go
	// 0.6931471805599453 1.4426950408889634 1e+09 8
}
