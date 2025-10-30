package main

import (
	"fmt"
	"reflect"
)

type S struct{ i int }

type Person struct {
	name string "namestr"
	age  int
}

func (p *S) Get() int  { return p.i }
func (p *S) Put(v int) { p.i = v }

func main() {
	fmt.Println("Hello")
}
