package main

import "fmt"

type User struct {
	name string
	age  uint8
}

func (user User) save() {
	fmt.Println("Usuário salvo", user.name, user.age)
}

func (user *User) changeAge() {
	user.age++
}

func main() {
	user := User{"George", 20}
	user.save()
	user.changeAge()
	fmt.Println(user.age)
}
