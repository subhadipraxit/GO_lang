package main

import (
	"fmt"
	s "strings"
)

var p = fmt.Println

func main() {

	p("String Operations : ")
	var str string = "New String"
	var strArray = []string{"This", "Separator,", "it?"}
	p(s.Contains(str, "New"))
	p(s.Repeat(str, 2))
	p(s.ReplaceAll(str, "New", "Old"))
	p(s.TrimSuffix(str, "ng"))
	p(s.Join(strArray, " is "))
	//p(s.s)  // for further try
	//p(s.)
	//p(s.)

}
