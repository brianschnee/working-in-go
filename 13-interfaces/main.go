package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

// value reciever
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// value reciever
func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}
	rectangle := Rectangle{width: 10, height: 5}

	fmt.Println(getArea(circle))
	fmt.Println(circle.area())
	fmt.Println(getArea(rectangle))
	fmt.Println(rectangle.area())
}
