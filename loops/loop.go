package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		time.Sleep(time.Second)
		fmt.Println("Incrementando: i")
		i++
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		time.Sleep(time.Second)
		fmt.Println("Incrementando: i")
	}

	nomes := []string{"George", "nadia", "brenda"}

	//o underline é o indice do laço
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for _, letra := range "OLA MUNDO" {
		fmt.Println(string(letra))
	}
}
