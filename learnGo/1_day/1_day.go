package main

import (
	"fmt"
	"unsafe"
)

const (
	m = 1 << iota
	n = 3 << iota
	o
	p
)

var str = "12345"

func main() {
	const (
		a = iota
		b
		c
		d = "str"
		e
		f = 100
		g
		h = iota
		i
	)

	println(a, b, c, d, e, f, g, h, i)
	println("-------------------------------")
	fmt.Println("m=", m)
	fmt.Println("n=", n)
	fmt.Println("n=", o)
	fmt.Println("p=", p)

	str = "12345"
	println(unsafe.Sizeof(str))

	var nub int = 10
	var pNub *int = &nub
	println(pNub)
	println(*pNub)
	println(&pNub)
	println(&nub)
}

//解释说明：ioat
/*
i=1：左移 0 位,不变仍为 1;
j=3：左移 1 位,变为二进制 110, 即 6;
k=3：左移 2 位,变为二进制 1100, 即 12;
l=3：左移 3 位,变为二进制 11000,即 24。
*/
