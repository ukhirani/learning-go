package main

// if you give the main function return type and also no parameters in the argument is allowed
// func main must have no arguments and no return values

func noReturnFunction(x int) {
	// this function is taking arguments but (ret1 type1) is the format of the function arguments
}

func ReturnFunction(x int) int {
	// we can also just say the return type of the function

	return 0
}

func returnFunctionWithVariable(x int) (hello int, yellow int) {
	return 7, 4
}

func Sum(a, b int) int { return a + b } // we can also write everything in the same line
