package main

import "fmt"

func main4() {
	type Grades int
	const (
		A Grades = iota
		B
		C
		D
		F
	)

	fmt.Printf("\nA: %d\nB: %d\nC: %d\nD: %d\nF: %d", A, B, C, D, F)
}
