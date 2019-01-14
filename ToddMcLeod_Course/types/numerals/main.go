package main

import "fmt"

/*
ox = hex notation denotation

we use base 10 numeral system
42420=
4=10^4 (4*10*10*10*10)
2=10^3 (2*10*10*10)
4=10^2 (4*10*10)
2=10^1 (2*10)
0=10^0 (0)

base 2 =
1 = 2^0 (2*0)
2 = 2^1 (2*1)
4 = 2^2 (2*2)
8 = 2^3 (2*2*2) or (2*2=4*2=8)
16 = 2^4 (2*2*2*2) or (2*2=4*2=8*2=16)

binary = base 2 in 0 and 1
42 =101010
32 = 2^5
8 = 2^3
2 = 2^1
911= 38f (3*16^2+8*16^1+f)

base 16 = 0-9 + ABCDEF
16^0 = ones
16^1 = 16's
16^2 = 256's
16^3 = 4096's
16^4 = 65536's

*/

func main() {

	s := "H"
	fmt.Println(s)

	//conver H to int
	bs := []byte(s)
	fmt.Println(bs) //H

	//assign the int from 45 to n
	n := bs[0]
	//print the value
	fmt.Println(n) //[72]
	//print the type
	fmt.Printf("%T\n", n) //72
	//print the binary translation
	//1 in the 64's position + 1 in the 8's position = 72
	fmt.Printf("%b\n", n) //1001000
	//print the hex conversion
	//4*16+8 = 72
	fmt.Printf("%#x\n", n) //0x48

}
