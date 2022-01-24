package main

import (
	"log"
	"net/http"
)

//principal lib para http -> net/http

func home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Seja bem vindo"))
}

func main() {
	http.HandleFunc("/home", home)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
