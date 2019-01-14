package main
//byte = alias for int8

import "fmt"

func main() {
	s := "Hello, playground"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	t := `Hello, playground`
	fmt.Println(t)
	fmt.Printf("%T\n", t)

	u := `"Hello, playground"`
	fmt.Println(u)
	fmt.Printf("%T\n", u)

	//slice of byte. convert s to int
	bs:=[]byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

}
