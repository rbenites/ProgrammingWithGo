package main

import "fmt"

//iota = pre-declared identifier
//

const (
	a = iota
	b = iota
	c = iota
	d = iota + 1
	e
)
const (
	f = iota
	g
	h
)

func main() {
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(c)
	fmt.Printf("%T\n", c)
	fmt.Println(d)
	fmt.Printf("%T\n", d)
	fmt.Println(e)
	fmt.Printf("%T\n", e)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
	fmt.Println(g)
	fmt.Printf("%T\n", g)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
}
