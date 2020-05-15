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

	// 2nd way to initilize struct filelds
	//	firstName, lastName, company string
	//	id                           int
}

// greeting method : value receiver
func (p Person) greet() string {

	return "My name is " + p.lastName + " , " + p.firstName + p.lastName + "," + strconv.Itoa(p.id) + " and I work for " + p.company

}

// change id : Pointer receiver
func (p *Person) chnageID() {
	p.id++
}

// switch company : Pointer receiver
func (p *Person) switchCompany(s string) {
	p.company = s
}
func main() {

	p := Person{"abc", "xyz", 333, "CGR"}                                          // 1st way to initialize
	var p1 = Person{"abcd", "ert", 222, "SA"}                                      // 2nd way to initialize
	var p2 = Person{firstName: "James", lastName: "Bond", id: 007, company: "Mi6"} // 3rd way to initialize
	pt(p)
	pt(p1)
	pt(p2)

	pt(p2.greet())

	p.chnageID()
	pt(p.greet())
	p.switchCompany("NYPD")
	pt(p.greet())

	q := Person{firstName: "QQ", lastName: "qq", id: 9, company: "Mi6"}

	pt(q.firstName, q.lastName, q.id)

	q_ptr := &q
	q_ptr.firstName = "JJ"

	pt(q)
}
