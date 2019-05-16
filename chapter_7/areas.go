package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt((a * a) + (b * b))
}

func (rec *Rectangle) area() float64 {
	l := distance(rec.x1, rec.y1, rec.x1, rec.y2)
	w := distance(rec.x1, rec.y1, rec.x2, rec.y1)
	return l * w
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	//	var rx1, ry1 float64   = 0, 0
	//	var rx2, ry2 float64   = 10, 10
	//	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	//	fmt.Println(circleArea(c))
	//      total := 0.0

	c := &Circle{0, 0, 5}
	r := &Rectangle{0, 0, 10, 10}

	// these are already pointers
	fmt.Println(totalArea(c, r))
}
