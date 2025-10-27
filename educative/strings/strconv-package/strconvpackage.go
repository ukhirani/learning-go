package main

import (
	"fmt"
	"strconv"
	str "strings"
)

func main() {
	my_string := "0"
	my_num := 0
	//for converting numbers to string
	int_converted_to_string := strconv.Itoa(my_num)

	fmt.Println(int_converted_to_string, str.TrimSpace(my_string))

	//there is also one for float
	fmt.Println(strconv.FormatFloat(3.14, 32, 3, 3)) // TODO: somehow find a way to make this work

}
