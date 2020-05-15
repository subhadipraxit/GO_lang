package main

import "fmt"

var pt = fmt.Println
var pf = fmt.Printf

type rect struct {
	height, width int
}

func (r rect) area() int {

	return r.height * r.width
}
func main() {

	rect1 := rect{4., 5}
	pt(rect1.area())
}
