// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

//конструктор структуры Point
func newPoint(a, b float64) Point{
	return Point{
		x: a,
		y: b,
	}
}

func main() {
	a := newPoint(1, 1)
	b := newPoint(2, 2)

	res := math.Sqrt(math.Pow(a.x - b.x, 2) + math.Pow(a.y - b.y, 2))

	fmt.Printf("Расстояние между X и Y %0.2f\n", res)
}