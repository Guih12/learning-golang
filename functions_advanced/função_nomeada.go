package main

import "fmt"

func calculate(n1, n2 int) (adc int, sub int) {
	adc = (n1 + n2)
	sub = (n1 - n2)
	return
}

func main() {
	adc, sub := calculate(50, 20)
	fmt.Println(adc, sub)
}
