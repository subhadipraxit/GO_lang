package main

import "fmt"
//global := 32 // Shorthand is not allowed outside function
func main(){

	// create variables using var keyword
	var name string = "Apple"
	var age int32 = 99
	const isCool = true

	global := 67 // Shorthand declaration
    //global + age // cannot add as two adifferent datatype
	//isCool = false // cannot assign value to const
	fmt.Println(name, age, isCool, global)

	// get the datatype of a variables
	fmt.Printf("The datatype for %d is %T\n", age,age)
}