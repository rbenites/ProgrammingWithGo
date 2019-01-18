package main

import "fmt"

/*If is a conditional statement
- bool: True or False
- not operator: !
initialization statement:
- if/else
- if/else if/else
- is/else if/else if/ ...else
*/
func main() {
	if true {
		fmt.Println("True")
	}
	if false {
		fmt.Println("False")
	}
	if !true {
		fmt.Println("Not True")
	}
	if !false {
		fmt.Println("Not False")
	}
	if x := 42; x == 42 {
		fmt.Println("x = 42")
	}
	fmt.Println("after x:=42")

	y := 42
	if y == 40 {
		fmt.Println("Y = 40")
	} else if y == 41 {
		fmt.Println("Y=41")
	} else {
		fmt.Println("Y=", y)
	}
//loop
	for i := 0; i < 100; i++ {
		//modulous with condition
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
