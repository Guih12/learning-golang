package main

import "fmt"

func sum(number int8, number_two int8) int8 {
	return number + number_two
}

//função que retorna mais de dois valores
func calculate(num1, num2 int8) (int8, int8) {
	sum := num1 + num2
	sub := num1 - num2
	return sum, sub
}

func main() {
	var soma int8 = sum(10, 30)
	println(soma)

	//parece arrow functions de javascript
	var function = func(word string) {
		fmt.Println(word)
	}

	function("George")

	result, _ := calculate(20, 30)
	fmt.Println(result)
}
