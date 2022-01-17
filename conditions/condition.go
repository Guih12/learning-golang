package main

import "fmt"

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
}
