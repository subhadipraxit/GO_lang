package main

import "fmt"

var pt = fmt.Println
var pf = fmt.Printf

func main() {

	arr := []int{33, 45, 76, 62, 12, 35, 39, 96}

	for i, id := range arr {
		pf("%d  -- %d \n", i, id)
	}

	for _, id := range arr {
		pf(" -- %d \n", id)
	}

	sum := 0

	for _, id := range arr {
		sum += +id
	}

	pt(sum)

	emails := map[string]string{"abc": "abc@email.com", "ijk": "ijk@emails.com", "xyz": "xyz@emails.com"}

	pt(emails)

	for k, v := range emails { // use/ print both key and values

		pf(" name : %s   - email : %s\n", k, v)
	}

	for k := range emails { // use/ print only key   , it is same as "for k,_:= range emails "  --> , _ is not required as it automaticall takes first variable as KEY

		pf("name : %s\n", k)
	}

	for _, v := range emails {

		pf("email : %s\n", v)
	}

}
