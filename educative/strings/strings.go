package main

import "fmt"

func main() {
	// strings are immuatbel in go and are of value type thus cannot be modified once declared
	hello := "string\n"      // this is an interpreted string
	hello_raw := `string \n` // this is a raw string and hence will also display /n
	fmt.Println(hello, hello_raw)
	// Note: Taking the address of a character in a string is illegal
	hello_combined := hello + hello_raw
	fmt.Println(hello_combined)

	hello += hello
	fmt.Println("new hello", hello)
}

//Contrary to strings in other languages as C++, Java or Python that are fixed-width (Java always uses 2 bytes), a Go string is a sequence of variable-width characters (each 1 to 4 bytes).
//go strings occupy less space because of variable width characters
