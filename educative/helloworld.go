package main

import "fmt"

// one new thing that i learnt that is in the same package you cannot redeclare
// so in the same package that main fucntion can reside in only one file/declaration
func main() {
	fmt.Println("Hello, World")
	// This illustrates the international character of Go by printing Καλημέρα κόσμε; or こんにちは 世界.
	fmt.Printf("Καλημέρα κόσμε; or こんにちは 世界\n")
	printMessage("Hello there from another functionn")
}

func printMessage(hello string) {
	fmt.Println(hello)
}
