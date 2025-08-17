package main

import "fmt"

func main11() {
	IdMap := map[string]int{
		"joe":  123,
		"jane": 456,
		"jack": 789,
	}

	fmt.Println(IdMap["jane"])
	IdMap["jane"] = 465
	fmt.Println(IdMap["jane"])

	for key, value := range IdMap {
		fmt.Println(key, value)
	}

	delete(IdMap, "jack")
	fmt.Println(IdMap["jack"])

	id1, p1 := IdMap["joe"]
	id2, p2 := IdMap["jack"]

	fmt.Println("Checking existance and value of Joe: ")
	fmt.Printf("%t\t%d\n", p1, id1)
	fmt.Println("Checking existance of Jack: ")
	fmt.Printf("%t\t%d\n", p2, id2)

	fmt.Printf("Length of Map: %d", len(IdMap))
}
