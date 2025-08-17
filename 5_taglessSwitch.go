package main

import "fmt"

func main5() {
	x := 1

	switch {
	case x > 1:
		fmt.Println("case 1")
	case x < 1:
		fmt.Println("case 2")
	default:
		fmt.Println("x=1")
	}
}
