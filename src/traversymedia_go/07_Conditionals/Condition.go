package main

import "fmt"

var prtf = fmt.Printf
var prt = fmt.Println

func main() {

	x := 8
	y := 8
	z := 91

	if x < y || x == y {
		prtf("%d is smaller than OR EQUAL to %d  \n", x, y)
	}

	// if-else and else-if
	if x < z && x+z == 19 {
		prt("Both the condition matched  \n")
	} else if x+z == 17 {
		prt("One condition matched  \n")
	} else {
		prt("No condition matched  \n")
	}

	color := "white"
	// Switch
	switch color {

	case "red":
		prt("Color is RED")
	case "white":
		prt("Color is WHITE")
	default:
		prt("Color is not RED or WHITE")
	}
}
