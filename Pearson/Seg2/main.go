package main

import (
	"fmt"
	"reflect"
)

var (
	name = "skshm"
	age int = 6
)

var globl string = "Global"

func main() {
	const what = "Something"
	var stronk int = 69
	bleep := "bloop"
	
	bobbl := testing(what)
	fmt.Println(bobbl)

	fmt.Println("Stronk and Bleep are " + reflect.TypeOf(stronk).String() + " and " + reflect.TypeOf(bleep).String())
	fmt.Println("Name and Age are " + reflect.TypeOf(name).String() + " and " + reflect.TypeOf(age).String())
}


func testing(what string,) string {
	fmt.Println("Hello, World! " + what + " is up!")
	return retStr(globl)
}

// Function that returns a string
func retStr(iny string) string {
	return "Yes!"
}