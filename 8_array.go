package main

import "fmt"

func main8() {
	var x [5]int
	x_literal := [...]int{1, 2, 3, 4, 5}

	x[0] = 1
	fmt.Printf("x[0] = %d\nx[1] = %d\n", x[0], x[1])

	for i, v := range x_literal {
		fmt.Printf("ind %d: val %d\n", i, v)
	}
}
