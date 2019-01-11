package main

import "fmt"

/*
- use var to DECLARE three variables (x,y,z)
- The variables should have package level scope (outside of func main)
- varivables should not have values assigedn to them
-  x= int, y= string, z= bool
- in func main print out the values for each identifier
*/

var x int
var y string
var z bool

func main() {
	fmt.Println(x) //0
	fmt.Println(y) //
	fmt.Println(z) //false
}
