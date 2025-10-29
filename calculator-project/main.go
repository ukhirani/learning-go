package main

import "fmt"

func main() {

	var operation string
	var num1, num2 int

	fmt.Println("UMANG'S CALCULATOR")
	fmt.Println("==================")

	fmt.Println("Please Enter the operation you want to perform (add,subtract,multiply,divide)")
	fmt.Scanf("%s", &operation)

	fmt.Println("Please enter both the numbers ")
	fmt.Scanf("%d %d", &num1, &num2)

	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "subtract":
		fmt.Println(num1 - num2)
	case "multiply":
		fmt.Println(num1 * num2)
	case "divide":
		fmt.Println(num1 / num2)
	default:
		fmt.Println("Invalid Operation")
	}
}
