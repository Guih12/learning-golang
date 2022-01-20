package main

import (
	"fmt"
	"time"
)

func writeDisplay(param string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- param
		time.Sleep(time.Second)
	}
	close(channel)
}

func main() {
	channel := make(chan string)
	go writeDisplay("Ola mundo", channel)
	for message := range channel {
		fmt.Println(message)
	}
}
