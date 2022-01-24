package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Dog struct {
	Name string `json:"nome"`
	Age  int    `json:"idade"`
}

func main() {
	dog := Dog{"Ratão", 2}
	convert, err := json.Marshal(dog) //converte em bytes
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(convert)
	fmt.Println(bytes.NewBuffer(convert)) //cria uma estrutura json em go
}
