package main

import (
	"fmt"
	"time"
)

func write(text string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- fmt.Sprintf("Valor recebido: %s", text)
			time.Sleep(time.Millisecond * 500)
		}
	}()
	return channel
}

func multiplexar(channelIpuntOne, channnelInputTwo <-chan string) <-chan string {
	channelOutput := make(chan string)

	for {
		select {
		case message := <-channelIpuntOne:
			channelOutput <- message
		case message := <-channnelInputTwo:
			channelOutput <- message
		}
	}
	return channelOutput
}

func main() {
	channel := multiplexar(write("Ola mundo"), write("Golang é legal"))

	for i := 0; i < 10; i++ {
		fmt.Println(channel)
	}
}
