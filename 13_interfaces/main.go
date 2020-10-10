package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	x, y, radius float64
}

type rectangle struct {
	w, h float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.h * r.w
}

func getArea(s shape) float64 {
	return s.area()
}

func main() {
	circle := circle{x: 0, y: 0, radius: 5}
	rectangle := rectangle{w: 10, h: 5}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))
}
