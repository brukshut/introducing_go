package main

import (
	"fmt"
	"math"
)

// types here
type Circle struct {
	x, y, r float64
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

type Shape interface {
	perimeter() float64
}

func totalPerimeter(shapes ...Shape) float64 {
	var totalPerimeter float64
	for _, s := range shapes {
		totalPerimeter += s.perimeter()
	}
	return totalPerimeter
}

// functions here
func (rec *Rectangle) perimeter() float64 {
	length := rec.x2 - rec.x1
	width := rec.y2 - rec.y1
	return 2 * (length + width)
}

func (c *Circle) perimeter() float64 {
	// actually, circumference
	return 2 * math.Pi * c.r
}

// end functions

// main
func main() {
	c := &Circle{0, 0, 5}
	r := &Rectangle{0, 0, 10, 10}
	fmt.Println(r.perimeter())
	fmt.Println(c.perimeter())
	// totalPerimeter() uses shapes interface
	fmt.Println(totalPerimeter(c, r))
}

// end main
