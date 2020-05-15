package main

import "fmt"

var pt = fmt.Println
var pf = fmt.Printf

func main() {

	// Long method
	i := 5
	for i >= 0 {
		pf("Long Loop counter %d\n", i)
		i--
	}

	// Short  method

	for i := 4; i >= 0; i-- {
		pf("Short Loop counter %d\n", i)
	}

	// FizzBuzz
	pt("Starting FizzBuzz")
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 {
			pt("FizzBuzz")
		} else if i%3 == 0 {
			pt("Fizz")
		} else if i%5 == 0 {
			pt("Buzz")
		} else {
			pt(i)
		}
	}
}
