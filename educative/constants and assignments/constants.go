package main

func getNumber() int {
	return 5
}

func main() {
	//A value that cannot be changed by the program is called a constant.
	//This data can only be of type boolean, number (integer, float, or complex) or string.

	// here's how we declare a constant in go
	//const identifier [type] = value

	const hello int = 5 //TODO: do all constants need to be used in GO like normal variables

	//Remark: There is a convention to name constant identifiers with all uppercase letters, e.g., const INCHTOCM = 2.54. This improves readability.
	// so my hello declaration is a NO NO (haha, got the joke ?)
	const PI = 3.14159

	//You may have noticed that we didn’t specify the type of constant PI here. It’s perfectly fine because the type specifier [type] is optional because the compiler can implicitly derive the type from the value.

	const B = "hello"
	//this is same as
	const BB string = "hello"

	//Typed and untyped constants
	//Constants declared through explicit typing are called typed constants, and constants declared through implicit typing are called untyped constants.

	// A value derived from an untyped constant becomes typed when it is used within a context that requires a typed value

	// 	var n int
	// f(n + 5)
	// untyped numeric constant "5" becomes typed as int, because n was int.

	//---------------------COMPILATION-----------------------------

	const C1 = 2 / 3       //okay
	const C2 = getNumber() //not okay (can you see the red squiggle)

	//WHY ????

	//Constants must be evaluated at compile-time. A const can be defined as a calculation, but all the values necessary for the calculation must be available at compile time.

	//Because the function getNumber() can’t provide the value at compile-time. A constant’s value should be known at compile time according to the design principles where the function’s value is computed at run time. So, it will give the build error: getNumber() used as value.

	// go run constants.go
	// # command-line-arguments
	// ./constants.go:38:13: getNumber() (value of type int) is not constant

	///fix getNumber() (value of type int) is not constant

	// To fix the error, replace the constant assignment with a value that is known at compile time instead of calling a function.

	// the above was said by chatgpt haha

}
