package main

import "fmt"

const (
	_ = iota
	//kb=1 shifted over ten times (1 with ten 0s)

	//1 shifted over ten spots
	kbb = 1 << (iota * 10)
	//1 shifted over 10^2 spots (20 spots)
	mbb = 1 << (iota * 10)
	//1 shifted over 10^3 spots (30 spots)
	gbb = 1 << (iota * 10)
)

func main() {
	x := 2
	//%d = decimal system
	//%b = binary system
	fmt.Printf("%d\t\t%b\n", x, x)

	//shifted the value over by one
	y := x << 1
	//shifted 0 from the 4s place to the 8s place
	fmt.Printf("%d\t\t%b\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	//shifted over ten spots
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	//shifted over ten spots
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t%b\n", kbb, kbb)
	fmt.Printf("%d\t\t%b\n", mbb, mbb)
	fmt.Printf("%d\t%b\n", gbb, gbb)
}
