package main

import (
	"errors"
	"fmt"
)

func main() {
	number := -100000000

	var number2 uint32 = 10000
	fmt.Println(number)
	fmt.Println(number2)

	//alias
	var number3 rune = 123456
	fmt.Println(number3)

	var number4 byte = 123
	fmt.Println(number4)

	number_real := 100000.25
	fmt.Println(number_real)

	//strings

	var str string = "Ola cara"
	str2 := "seja legal"
	fmt.Println(str)
	fmt.Println(str2)

	//errors
	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
