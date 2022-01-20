package main

import (
	"fmt"
	"time"
)

func writeDisplay(param *string) {
	for {
		fmt.Println(*param)
		time.Sleep(time.Second)
	}
}

func main() {
	//concorrência != paralelismo
	param := "Seja bem vindo"
	go writeDisplay(&param)
	param2 := "Programando em go"
	writeDisplay(&param2)
}
