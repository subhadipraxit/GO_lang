package main

import "fmt"

var pt = fmt.Println
var pf = fmt.Printf

func main() {

	emails := make(map[string]string) // Type one declaration

	emails["abc"] = "abc@email.com"
	emails["xyz"] = "xyz@email.com"
	emails["ijk"] = "ijk@email.com"

	pt(emails)
	pt(len(emails))

	delete(emails, "xyz") // How to delete elements
	pt(emails)

	altEemails := map[string]string{"abc": "cba@email.com"} // Type two declaration

	pt(altEemails)
}
