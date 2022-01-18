package main

import "fmt"

func one() {
	fmt.Println("Executed one")
}

func two() {
	fmt.Println("Executed two")
}

func main() {
	defer one()
	two()
}
