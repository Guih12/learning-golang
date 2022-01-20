package main

import (
	"fmt"
	"sync"
	"time"
)

func writeDisplay(param string) {
	for i := 0; i < 5; i++ {
		fmt.Println(param)
		time.Sleep(time.Second)
	}
}

func main() {
	//concorrência != paralelismo
	var waitGroup sync.WaitGroup

	waitGroup.Add(4) //grupo de espera e tem duas go routines

	go func() {
		writeDisplay("Excutando go goutine 1")
		waitGroup.Done() //ele tira um do contador de go routines
	}()

	go func() {
		writeDisplay("Excutando go goutine 2")
		waitGroup.Done()
	}()

	go func() {
		writeDisplay("Excutando go goutine 3")
		waitGroup.Done()
	}()

	go func() {
		writeDisplay("Excutando go goutine 4")
		waitGroup.Done()
	}()

	waitGroup.Wait() //fala pra esperar a quantidade ser == 0 para terminar
}
