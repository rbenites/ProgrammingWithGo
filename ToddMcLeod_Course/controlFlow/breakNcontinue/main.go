package main

import "fmt"

func main() {
	x := 1
	for {
		//increment x by 1
		x++
		if x > 100 {
			//if x is greater than 100 break out
			break
		}
		//if x has a remainder continue (loop back up) but if it does not have a remainder print the value
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	//once x is greater than 100 print done
	fmt.Println("done")
}
