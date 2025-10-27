package main

import (
	"fmt"
	"strings"
)

func main() {
	hello := "umangum"

	prefix_result := strings.HasPrefix(hello, "um")
	suffix_result := strings.HasSuffix(hello, "um")
	substring_result := strings.Contains(hello, "um")
	substring_index := strings.Index(hello, "um") //else it will return -1 if it doesn't exist
	last_index := strings.LastIndex(hello, "um")  //else it will return -1 if it doesn't exist
	//strings.IndexRune(s string, ch int) int --> if you are searching for something non ascii
	if prefix_result {
		fmt.Println("prefix")
	}

	if suffix_result {
		fmt.Println("suffix")
	}

	if substring_result {
		fmt.Println("substring")
	}

	fmt.Println(substring_index)
	fmt.Println(last_index)

	//replacing substring

	myname := "Umang Hirani"
	var newname string

	fmt.Println("Before replacing :", myname)

	newname = strings.Replace(myname, "Umang", "The Umang", 1) // here the count n -> the last arg means the no. of occurences from start that will be replaced
	// if we have -1 as the last arg -> then all will be replaced
	//
	// instead we can use the below for replacing all occurences

	newname = strings.ReplaceAll(newname, "Hirani", "The Great Hirani")
	fmt.Println("After replacing :", newname)

	//here's how to count the occurences of a substring in a string
	str_count := strings.Count(newname, "The")
	fmt.Println("Here is the count :", str_count)

	//here's how to repeat a string 'n' times
	repeated_string := strings.Repeat(newname, 3)
	fmt.Println(repeated_string)

	// toUpper and toLower
	fmt.Println("Here's the lower case version", strings.ToLower(newname))
	fmt.Println("Here's the upper case version", strings.ToUpper(newname))

	//Trimming is also available
	very_bad_string := "    asdfasdf asdfasdf    "
	fmt.Println(strings.TrimSpace(very_bad_string))

	//you can also remove leasing and trailing from the string
	fmt.Println(strings.Trim(very_bad_string, " "))                    // now it's the same as the above function
	fmt.Println(strings.Trim(strings.TrimSpace(very_bad_string), "a")) // now it's the same as the above function

	//============================  splitting a string  =============================================================

	mystring := "Umang Ketanbhai Hirani"
	//on whitespaces
	fmt.Println(len(strings.Fields(mystring))) // here you won't see different slices due to Println function

	// on a separator
	mystring = "Umang:Ketanbhai:Hirani"
	fmt.Println(len(strings.Split(mystring, ":")))
	split_slice := strings.Split(mystring, ":")

	fmt.Println(strings.Join(split_slice, " THE ")) // joins the splice to a string with a delimiter that you mentioned

	// Reading from a string
	/*The strings package also has a function called strings.NewReader(str). This produces a pointer to a Reader value, that provides amongst others the following functions to operate on str:

	Read() to read a []byte
	ReadByte() to read the next byte from the string.
	ReadRune() to read the next rune from the string.
	*/

	//conversion to and from a string
}
