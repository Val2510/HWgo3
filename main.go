package main

import (
	"fmt"
	"math"
)

type Shape struct {
	Name string
}

func (s Shape) GetName() string {
	return s.Name
}

func (s Shape) Area() float64 {
	return 0.0
}

type Rectangle struct {
	Shape
	width  float64
	height float64
}

type Circle struct {
	Shape
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.height * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	rec := Rectangle{Shape: Shape{Name: "Rectangle"}, width: 8, height: 6}
	circ := Circle{Shape: Shape{Name: "Circle"}, radius: 4}
	fmt.Println(rec.GetName(), "=", rec.Area())
	fmt.Println(circ.GetName(), "=", circ.Area())
}
