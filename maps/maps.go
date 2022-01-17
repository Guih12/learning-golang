package main

import "fmt"

func main() {
	fmt.Println("Maps")
	user := map[string]string{
		"nome":     "george",
		"lastname": "borsato",
	}

	fmt.Println(user["nome"])

	user_two := map[string]map[string]string{
		"user": {
			"primeiro": "Joao",
			"curso":    "Alves",
		},
	}

	fmt.Println(user_two)

}
