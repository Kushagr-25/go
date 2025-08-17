package main

import (
	"fmt"
	"strings"
)

func main7() {
	var x string
	_, err := fmt.Scan(&x)

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	x_lower := strings.ToLower(x)

	if strings.HasPrefix(x_lower, "i") && strings.HasSuffix(x_lower, "n") && strings.Contains(x_lower, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
