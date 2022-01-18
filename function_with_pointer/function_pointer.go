package main

import "fmt"

func invert(number int) int {
	return number * -1
}

func invertWithPointer(number *int) {
	*number = (*number * -1)
}

func main() {
	//number := 20
	//inv := invert(number)
	newNumber := 50
	fmt.Println(newNumber)
	invertWithPointer(&newNumber)
	fmt.Println(newNumber)
}
