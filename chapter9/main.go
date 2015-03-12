package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {

}

func main() {
	c := Circle{0, 0, 5}
	// fmt.Println(circleArea(&c))
	fmt.Println(c.area())
}
