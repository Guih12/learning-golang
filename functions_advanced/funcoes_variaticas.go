package main

import "fmt"

func soma(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

//podemos passar nos parâmetros um slice de numeros e quaisquer outro parãmetro
func show(word string, numbers ...int) (app string, result int) {
	app = word
	result = 0
	for _, number := range numbers {
		result += number
	}
	return
}

func main() {
	result := soma(5, 10, 20, 30)
	fmt.Println(result)
	fmt.Println(show("Go é maneiro", 20, 30, 40, 50, 60))
}
