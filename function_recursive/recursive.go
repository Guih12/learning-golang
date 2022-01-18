package main

import "fmt"

func fibonnaci(position int) uint {
	if position <= 1 {
		return uint(position)
	}
	return fibonnaci(position-2) + fibonnaci(position-1)
}

func main() {
	position := uint(10)

	for i := uint(0); i < position; i++ {
		fmt.Println(fibonnaci(int(i)))
	}
}
