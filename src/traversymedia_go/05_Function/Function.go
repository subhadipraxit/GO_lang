package main

import "fmt"

var p = fmt.Println

func add(x int, y int) int { // method 1 to declare  and return

	return x + y
}

func subtract(x, y int) (op int) { // method 2 to declare  and return

	op = x - y
	return
}

func calc(x, y int) (op1, op2, op3, op4 int) { // method 3 to declare and return

	op1 = x + y
	op2 = x - y
	op3 = x * y
	op4 = x / y
	return
}

func greeting(name string) string { // same as method1

	return "Hello" + name
}

func main() {

	p(add(4, 3))
	p(subtract(4, 3))
	p(calc(6, 3))

	p(greeting(" GO"))

}
