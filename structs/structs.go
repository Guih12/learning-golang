package main

import "fmt"

//tipo usuário
type Person struct {
	name    string
	age     uint8
	address Address
}

type Address struct {
	city   string
	number int
}

func main() {
	var u Person
	u.name = "George"
	u.age = 20

	fmt.Println(u)
	address := Address{"Rua Urias ribeiro", 12}

	user := Person{name: "Guilherme", age: 35, address: address}
	fmt.Println(user)

	user_two := Person{name: "Lemos"}
	fmt.Println(user_two)

}
