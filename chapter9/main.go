package main

import "fmt"

type Circle struct {
	x, y, r float64
}

func main() {
	// var c Circle
	// c := new(Circle)
	// c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	fmt.Println(c)
}
