package main

import "fmt"

func main() {
	var choice int

	fmt.Println("Choose a function to run: ")
	fmt.Println("1. hello.go")
	fmt.Println("2. varScope.go")
	fmt.Println("3. typeConversion.go")
	fmt.Println("4. iota.go")
	fmt.Println("5. taglessSwitch.go")
	fmt.Println("6. Truncation.go")
	fmt.Println("7. Findian.go")
	fmt.Println("8. Array.go")
	fmt.Println("9. Slices.go")
	fmt.Println("10. Slice Functions.go")
	fmt.Println("11. Maps.go")
	fmt.Println("12. Structs.go")
	fmt.Println("13. Struct Question.go")
	fmt.Println("Enter your choice (1-13): ")

	_, err := fmt.Scan(&choice)

	if err != nil {
		fmt.Println("Invalid Input")
		return
	}

	switch choice {
	case 1:
		fmt.Println("Calling hello.go:")
		main1()
	case 2:
		fmt.Println("Calling var_scope.go:")
		main2()
	case 3:
		fmt.Println("Calling typeConversion.go")
		main3()
	case 4:
		fmt.Println("Calling iota.go")
		main4()
	case 5:
		fmt.Println("Calling taglessSwitch.go")
		main5()
	case 6:
		fmt.Println("Calling trunc.go")
		main6()
	case 7:
		fmt.Println("Calling findian.go")
		main7()
	case 8:
		fmt.Println("Calling array.go")
		main8()
	case 9:
		fmt.Println("Calling slices.go")
		main9()
	case 10:
		fmt.Println("Calling sliceFunc.go")
		main10()
	case 11:
		fmt.Println("Calling maps.go")
		main11()
	case 12:
		fmt.Println("Calling struct.go")
		main12()
	case 13:
		fmt.Println("Calling structQuest.go")
		main13()
	default:
		fmt.Println("Invalid Choice")
	}
}
