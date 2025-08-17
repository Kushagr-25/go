package main

import "fmt"

func iterate(slice []int) {
	for i, v := range slice {
		fmt.Printf("ind %d: val %d\n", i, v)
	}

	fmt.Printf("Length of Slice = %d\nCapacity of Slice = %d\n", len(slice), cap(slice))
}
func main9() {
	arr := [...]int{1, 2, 3, 4, 5, 6}

	sli1 := arr[1:3]
	sli2 := arr[2:5]

	fmt.Println("Iterating slice 1: ")
	iterate(sli1)

	fmt.Println("iterating slice 2: ")
	iterate(sli2)
}
