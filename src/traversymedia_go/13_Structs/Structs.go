package main

import (
	"fmt"
	"strconv"
)

var pt = fmt.Println
var pf = fmt.Printf

type Person struct {
	firstName string
	lastName  string
	id        int
	company   string
}

func (p Person) greet() string {

	return "My name is " + p.lastName + " , " + p.firstName + p.lastName + "," + strconv.Itoa(p.id) + " and I work for " + p.company

}
func main() {

	p := Person{"abc", "xyz", 333, "CGR"}                                          // 1st way to initialize
	var p1 = Person{"abcd", "ert", 222, "SA"}                                      // 2nd way to initialize
	var p2 = Person{firstName: "James", lastName: "Bond", id: 007, company: "Mi6"} // 3rd way to initialize
	pt(p)
	pt(p1)
	pt(p2)

	pt(p2.greet())

}
