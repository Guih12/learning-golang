package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	//slice -> o slice sempre referĂȘncia um outro array
	slice := []int{10, 20, 30, 40, 50}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array2))

	slice = append(slice, 90)
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "90"
	fmt.Println(slice2)

	//arrays internos

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(cap(slice3))
}
