package main

import (
	"fmt"
	"math"
)

var pf = fmt.Printf
var pt = fmt.Println

type geometry interface {
	area() float64
	perim() float64
}

type circle struct {
	radius float64
}

type rect struct {
	height, width float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func (r rect) area() float64 {
	return r.height * r.width
}

func (r rect) perim() float64 {
	return 2*r.height + r.width
}

func measure(g geometry) {

	pf("Measurement  of %T", g)
	pt(g)
	pt("Area is : ", g.area())
	pt("Perimeter is : ", g.perim())

}
func main() {

	c1 := circle{radius: 4}
	c2 := circle{6.34}

	r1 := rect{height: 3, width: 2}
	r2 := rect{7.8, 5.2}

	measure(c1)
	measure(c2)

	measure(r1)
	measure(r2)

}
