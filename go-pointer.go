package main

import "fmt"

func main() {
	x := 15
	// a points to the memory address of x
	a := &x

	// Memory address - &
	// Value at memory address - *

	fmt.Println(a)
	fmt.Println(*a)
	// *a evaluates to variable x
	*a = 5
	fmt.Println(x)
	*a = *a * *a
	fmt.Println(x)
	fmt.Println(*a)
}
