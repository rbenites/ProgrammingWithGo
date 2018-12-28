package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	//give jimPointer the memory address of the jim variable
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()
}

//pointer to person receives the memory address from jimPointer. *person points to person.
func (pointerToPerson *person) updateName(newFirstName string) {
	//*pointerToPerson changes the address to the value in peron.firstname and changes it to the name passed from the update function call
	(*pointerToPerson).firstName = newFirstName
}

// call the function on any type person and within the function refrence to it with p
func (p person) print() {
	fmt.Printf("%+v", p)
}
