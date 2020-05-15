package main

import "fmt"

var pt = fmt.Println
var pf = fmt.Printf

func main() {

	a := 9
	b := &a
	pt(a, b)

	pf("Type of  variable b is %T\n", b)
	pf("printing the value of  the address which is stored in b is %d\n", *b)

	pf("Printing the value of a : %d\n", *&a) // *&a  is same as a

	// change value through pointer

	*b = 10

	pt(a)
}
