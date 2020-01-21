package add

import (
	"fmt"
)

type Number struct {
	I int
}

var num1 int = 11

var Num1 int = 10

/*
在go中，函数名称 以及全局变量名称首字母大小写代表着函数是共有和私有，
大写代表共有，小写代表私有。
*/
func Add(nub1 int, nub2 int) (ret int) {
	return nub1 + nub2
}

/*
函数接口的使用
多肽方式
*/
func Call() {
	var phone Phone

	phone = new(NokiaPhone)
	phone.call()

	phone = new(Iphone)
	phone.call()
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (nokia NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you!")
}

type Iphone struct {
}

func (iphone Iphone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func init() {
	fmt.Println("add init call")
	Sub(3, 1)
}
