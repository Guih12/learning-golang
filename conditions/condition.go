package main

import "fmt"

func dayOfWeek(number int) string {
	switch number {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	default:
		return "Valor invalido"
	}
}

//outra forma de escrever switch com go
func dayOfWeekTwo(number int) string {
	switch {
	case number == 1:
		return "Domingo"
	default:
		return "Valor invalido"
	}
}

func main() {

	number := 10
	if number > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("É menor que 15")
	}

	//IF INIT
	if otherNumber := number; otherNumber > 0 {
		fmt.Println("numero é maior que zero")
	}

	dia := dayOfWeek(1)
	fmt.Println(dia)

	dia2 := dayOfWeekTwo(1)
	fmt.Println(dia2)
}
