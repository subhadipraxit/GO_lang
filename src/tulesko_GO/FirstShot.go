package main

import (
	"fmt"
	"math"
)
var m = 88

func main() {

	fmt.Println("This is first shot")

	var num = 5  // declared without datatype
	var num1 int = 2  // declared as integer
	var num2 uint = 3 // declared as unsigned integer
	var number int //cannot declare without datatype while we don't assign values
	const constant = 10 // declared as constant
//	constant = 8 // cannot change the value as declared as constant(const)
	number = 7 
	number1 := 9 // it's short hand of --> var number1 = 9
   //	var result = num1 + num2 // we cannot add two different datatype
	    var result = num1+ number
 		 fmt.Println(result)
		 fmt.Println(num) // we cannot keep unused variables in code
		 fmt.Println(num2) // we cannot keep unused variables in code
		 fmt.Println(number1)
		 fmt.Println(3+5)

	var num3 , num4 int

	fmt.Println(num3+num4) // Print with default value
	 
	var num5 int32=5 
	
	var num6 int32= 6
	
	fmt.Println(num5+num6)

	//for {
	//   fmt.Println("It is infinite loop")
	//    }
	
	//Method1 for loop
	i := 1
	for i <=5 {
		   fmt.Println("Go for loop : " , i)
		   i++
	    }
	//Method2 for loop	
		for j:=10;j>5;j-- {
			fmt.Println("Go for loop : " , j)
		 }

	// Call function

	a :=10
	b := 5

	fmt.Println(add(a,b))
	fmt.Println(subs(a,b))
	res1, res2 := calc(a,b) // assign/catch two values simultaneously
	fmt.Println(res1,res2)

	demo()
	fmt.Println(m)

	fmt.Println(math.Sqrt(17))
	var ans = math.Floor(math.Sqrt(17))
	var ans1 = math.Ceil(math.Sqrt(17))

	 ans = math.Round(ans)
	fmt.Println(ans)
	fmt.Println(ans1)

	var ans2= math.Sqrt(17)
	fmt.Printf("The square root result is  %f  or %.2f by compressing, Thanks", ans2, ans)
	fmt.Printf("The square root result is  %d , Thanks", m)
	
	
	
	
	fmt.Println()
  
}


func demo(){

	m := 8
	fmt.Println(m)
}

func add(x int, y int ) int{

	 z :=x+y
	 return z
}

func subs(x , y int ) int{

	z :=x-y
	return z
}

func calc(x , y int ) (int,int){  // retun multiple values

	z1 := x*y
	z2 := x/y
	return z1, z2
}

func calculator(x , y int ) (out1, out2 int){  // 2nd way retun multiple values

	out1 = x*y
	out2 = x/y
	return
}