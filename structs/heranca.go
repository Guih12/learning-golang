package main

import "fmt"

//modelo de classe abstrata
type Person struct {
	name     string
	lastname string
	age      uint8
	height   float32
}

//modelo de herança
type Study struct {
	Person
	curso      string
	university string
}

func main() {
	person := Person{"george", "borsato", 21, 1.66}
	study := Study{person, "Ads", "UFMS"}
	fmt.Println(study)
}
