package main

import "fmt"

func find(slice []int) {
	fmt.Printf("Length = %d\nCapacity = %d\n", len(slice), cap(slice))
}

func main10() {
	sli1 := make([]int, 1)
	sli2 := make([]int, 0, 2)

	fmt.Println("For Slice 1:")
	find(sli1)

	fmt.Println("For Slice 2:")
	find(sli2)

	sli2 = append(sli2, 2, 5, 8, 10)

	fmt.Println("Slice 2 after appending: ")
	iterate(sli2)
}
