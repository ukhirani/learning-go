package main
import "fmt"

//this is basically how you import anything in go 
//also the first line means the package that it belongs to 
//fmt means formatted IO

import "fmt"
import "os"

//the above is also ok and the below is also ok 

import "fmt"; import "os"

//but when you are importing multiple packages,it's just best to write it as below :

import (
  "fmt"
  "os"
)

//this above thing is known as factoring which means calling the keyword once on multiple instances
//
//It is also applicable to keywords like const, var, and type.
//

int hello // this is the same as calling it private and it will be visible only inside the package.
int Hello // this is the same as calling it public and it will be visible even outside the package also.main

//however this is how you would give an alias to a package if you wanted to 

import fm "fmt"
//here the package fmt is imported as fm (same as how we would do it in python)

func (){
	fmt.func1();
	fmt2.func1();
	//here both are called but you see they are reffering to different objects as they are not of the same package my dear
}

//Note: Go has a motto known as “No unnecessary code!”. So importing a package which is not used in the rest of the code is a build-error.


