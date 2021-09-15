package main

import (
	"fmt"
	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	/*a := 1>>2
	b := -8>>2
	c := 1<<2
	d := -2<<2
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)*/
	fmt.Println(2&3)
	fmt.Println(2|3)
	fmt.Println(5&4)
	fmt.Println(-3|3) //-1
	fmt.Println(-3^3) //-2
	/*
	1111 1101
	0000 0011
 |:	1111 1111 (计算机中存储补码，有符号数，故-1)
 ^: 1111 1101 (计算机中存储补码，有符号数，故-2)
	 */
}
