package main

import "fmt"

func apresenta(channel chan string) {
	for message := range channel {
		fmt.Println(message)
	}
}

func main() {
	channel := make(chan string, 2) //criando um canal com buffer com o valor máximo de 2
	channel <- "Seja bem vindo"
	channel <- "aaa"
	close(channel)

	apresenta(channel)
}
