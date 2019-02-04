package main

import "fmt"

/*
- Create new type: VEHICLE
	- underlying type is a struct
	- The fields should have:
		- Doors
		- Color
- Create two new types: TRUCK & SEDAN
	- The underlying type of each of these new type is a struct
	- Embed the "vehicle" type in both truck & sedan
	- Give truck the field "fourWheel" which will be set to bool.
	- Give the dedan the field "luxury" which will be set to bool.
- Using the VEHICLE, TRUCK, and SEDAN structs
	- use a comparison literal, create a value of type truck and assign values to the fields
	-Using composite literal, create a value of type sedan and assign values to the fields.

*/

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourwheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourwheel: true,
	}
	fmt.Println(t)

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println(s)
}
