package main

import "fmt"

func main6() {
	var x float64
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	x_trunc := int(x)
	fmt.Println(x_trunc)
}
