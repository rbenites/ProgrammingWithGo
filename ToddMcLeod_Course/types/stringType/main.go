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
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	//converts each character into UTF8
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	fmt.Println("")

	//print index  position and the int conversion of the string
	for i, v := range s {
		fmt.Println(i, v)
	}

	//print the index position of the string and the hex decimal
	for i, v := range s {
		fmt.Printf("at index position %d we have %#x\n", i, v)
	}
}
