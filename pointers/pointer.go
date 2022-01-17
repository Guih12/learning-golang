package main

import "fmt"

func main() {
	fmt.Println("Ponteirio")

	var number int = 10
	var number_two = number
	fmt.Println(number, number_two)

	number_two++
	fmt.Println(number, number_two)

	//ponteiro é refêrencia de memorial
	var number_three int
	var pointer *int

	number_three = 100
	pointer = &number_three

	fmt.Println(number_three, pointer)

	//desreferenciação do ponteiro
	fmt.Println(number_three, *pointer)

	number_three = 150
	fmt.Println(number_three, pointer)
	fmt.Println(number_three, *pointer)
}
