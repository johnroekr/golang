package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	var s Shape

	c := Circle{10.}
	s = c
	fmt.Println("- Circle")
	fmt.Println("area", s.area())
	fmt.Println("perimeter", s.perimeter())

	r := Rect{5., 4.}
	s = r
	fmt.Println("- Rect")
	fmt.Println("area", s.area())
	fmt.Println("perimeter", s.perimeter()
}
