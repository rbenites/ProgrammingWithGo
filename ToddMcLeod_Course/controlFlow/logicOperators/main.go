package main

import "fmt"

func main() {
	fmt.Printf("true && true\t %v\n", true && true)   //true and true both must match
	fmt.Printf("true && false\t %v\n", true && false) // true and false, both must match
	fmt.Printf("true or true\t %v\n", true || true)   //true or true either must match
	fmt.Printf("true or false\t %v\n", true || false) //true or false, either must match
	fmt.Printf("not true\t %v\n", !true)              //Not true
	fmt.Printf("not false\t %v\n", !false)            //Not flase

}
