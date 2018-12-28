//package is a collection of common source code files. 1st line must define the file it belongs to
package main

//gives our package access to code in another package.
import "fmt" //this is a standard go library

//funciton with its name and an argument list
func main() {
	//tells go to access fmt.Printlin and print Hi there to the terminal
	fmt.Println("Hi there!")
}
