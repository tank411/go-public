package main

import (
	"fmt"
)

const port = "6789"

func main() {
	fmt.Println("runing main")
	go show()
	ret := add(12, 32)
	fmt.Println(ret)
}
func show() {
	for index := 0; index < 1000; index++ {
		fmt.Println("show", index)
	}
}

func add(a int, b int) (c int) {
	return a + b
}
