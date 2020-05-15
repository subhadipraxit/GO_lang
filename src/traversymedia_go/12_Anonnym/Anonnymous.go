package main

import "fmt"

var pf = fmt.Printf
var pt = fmt.Println

func seq() func(i int) int {
	sum := 0
	return func(x int) int {

		sum += x
		return sum
	}
}

func main() {

	// 1st example of anonymous function

	func(s string) {
		pt("Hello ", s)
	}("World")

	// 2nd example of anonymous function

	pt(func(x, y int) int {
		return x + y
	}(4, 5))

	// Closure

	sequence := seq()
	pt(sequence(5), sequence(5))
	pt(sequence(5))
}
