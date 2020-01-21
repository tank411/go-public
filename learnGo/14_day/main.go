/*
 * @Descripttion: 模块信息介绍
 * @Company: HIK
 * @version: 1.0.0
 * @Author: TK
 * @Date: 2019-06-25 16:23:37
 * @LastEditors: TK
 * @LastEditTime: 2019-12-04 17:06:29
 */
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
