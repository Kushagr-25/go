package main

import "fmt"

func main3() {
	var x int32 = 1
	var y int16 = 2

	fmt.Printf("\nBefore Conversion\nx: %d\ty: %d\n", x, y)

	x = int32(y)
	fmt.Printf("After Conversion\nx: %d\ty: %d", x, y)
}
