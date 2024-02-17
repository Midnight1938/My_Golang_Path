package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int = 3
	if a == 3 {
		fmt.Println("This is true")
	} else if a != 4 {
		fmt.Println("This is also true")
	} else {
		fmt.Println("This is false")
	}

	nummers := []int{7, 6}
	switchy(nummers)

}

func switchy(a []int) {
	ans := dubblInts(a[0], a[1]) % 2 == 0

	switch ans {
	case ans:
		fmt.Println("An Even " + strconv.Itoa(dubblInts(a[0], a[1])))
	default:
		fmt.Println("Other")
	}
}

func dubblInts(a, b int) int {
	return a * b
}
