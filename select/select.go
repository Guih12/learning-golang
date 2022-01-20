package main

import (
	"fmt"
	"time"
)

func main() {
	channelOne, channelTwo := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channelOne <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Millisecond * 2)
			channelTwo <- "Canal 2"
		}
	}()

	for {

		select {
		case messageChannelOne := <-channelOne:
			fmt.Println(messageChannelOne)
		case messagrChannelTwo := <-channelTwo:
			fmt.Println(messagrChannelTwo)
		}
	}
}
