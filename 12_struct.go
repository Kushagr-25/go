package main

import "fmt"

type Person struct {
	name string
	addr string
	phn  string
}

func printStruct(p *Person) {
	fmt.Println(p.name, p.addr, p.phn)
}

func main12() {
	//Assignment 1
	var p1 Person
	p1.name = "Joe"
	p1.addr = "1st Cross"
	p1.phn = "1234"

	//Assignment 2
	p2 := new(Person)
	p2.name = "Jane"
	p2.addr = "Church street"
	p2.phn = "46789"

	//Assignment 3
	p3 := Person{
		name: "Jack",
		addr: "King Road",
		phn:  "826",
	}

	printStruct(&p1)
	printStruct(p2)
	printStruct(&p3)

}
