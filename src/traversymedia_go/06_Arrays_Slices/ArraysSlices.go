package main

import "fmt"

var p = fmt.Println

func main() {

	var arr = [2]string{"abc", "fox"}
	arr1 := [6]string{"abc", "fox"}
	arr2 := []string{"abc", "fox", "times", "yorker"} // called slices

	var fruitArray [3]string

	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"
	fruitArray[2] = "Dragon"

	p(arr, arr1, arr2, fruitArray, fruitArray[2])
	p(len(arr2)) // print length
	p(arr2[1:4]) // print values between "1st index"   and "before 4th index"

}
