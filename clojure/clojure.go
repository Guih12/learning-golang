package main

import "fmt"

func clojure() func() {
	text := "Seja bem vindo"
	function := func() {
		fmt.Println(text)
	}

	return function
}

func main() {
	text := "Ola"
	fmt.Println(text)

	result := clojure()
	result()
}
