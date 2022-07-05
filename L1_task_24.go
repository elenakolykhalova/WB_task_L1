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
	var a, b float64
	fmt.Scan(&a, &b)
	c := newPoint(a, b)
	fmt.Printf("Расстояние между X и Y %0.2f\n", math.Abs(c.x - c.y))
}