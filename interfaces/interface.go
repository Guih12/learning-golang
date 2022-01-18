package main

import (
	"fmt"
	"math"
)

type forma interface {
	area() float64
}

type retangulo struct {
	width  float64
	height float64
}

type circulo struct {
	raio float64
}

func writeArea(f forma) {
	fmt.Println(f.area())
}

func (r retangulo) area() float64 {
	return r.height * r.width
}

func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	retangulo := retangulo{20, 30}
	writeArea(retangulo)

	circulo := circulo{20}
	writeArea(circulo)
}
