package main

import (
	"fmt"
	"math"
)

/*
func function_name( [parameter list] ) [return_types] {
   函数体
}


申明函数的标志func 函数名function_name，参数列表parameter list，返回值类型return_types，
go里面支持返回多个值，返回值对应就行。
*/
/*
说明：和c++相比，函数可以写在调用的后面，在c中函数写在调用的后面，编译器会提示函数未定义的错误。
在go里面可以允许写在后面，不用声明。

函数的参数使用默认的值传递，
*/

func add(a, b int) int {
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	return a + b
}

func main() {
	var marks string = "11"

	switch marks {
	case "10":
		fmt.Printf("10\n")
	case "11":
		fmt.Printf("11\n")
	default:
		fmt.Printf("marks default\n")
	}

	var grade string = "B"
	switch {
	case grade == "A":
		fmt.Printf("A\n")
	case grade == "B":
		fmt.Printf("B\n")
	default:
		fmt.Printf("grade default\n")
	}

	var typeName int = 3

	switch typeName {
	case 1, 2, 3:
		fmt.Printf("1或者2或者3\n")
	default:
		fmt.Printf("typeName default\n")
	}

	fmt.Printf("-------------------------------------\n")
	//和c的遍历方法一样
	for i := 0; i < 10; i++ {
		println(i)
	}
	fmt.Printf("-------------------------------------\n")
	//

	var a int
	var b int = 10
	for a < b {
		a++
		fmt.Printf("a 的值为：%d\n", a)
	}
	fmt.Printf("-------------------------------------\n")

	numbers := [3]int{1, 2, 3}

	for i, a := range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, a)
	}

	//------------------------------------------------
	var a1 int = 10
	var b1 int = 20
	add(a1, b1)
	fmt.Printf("-------------------------------------\n")

	a2, b2 := swap("a1", "b1")
	fmt.Println(a2, b2)
	fmt.Printf("-------------------------------------\n")

	//-----------------go 语言中的指针使用--------------------

	var a3 int = 10
	var b3, c3 = &a3, &a3
	fmt.Println(b3, c3)
	fmt.Println(&b3, &c3)
	fmt.Println(*b3, *c3)
	fmt.Println(a3)

	*b3 = *b3 + 10

	fmt.Println(b3, c3)
	fmt.Println(&b3, &c3)
	fmt.Println(*b3, *c3)
	fmt.Println(a3)

	fmt.Printf("-------------------------------------\n")

	//-----------go 中函数作为值使用------------------
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	/* 使用函数 */
	fmt.Println(getSquareRoot(9))
	fmt.Printf("--------------函数作为值使用-----------------------\n")

	//----------------闭包的使用---------------------
	var a4 int = 5
	fun1 := func() func() {
		var i int = 10
		return func() {
			fmt.Println("i,a5:%d,%d\n", i, a4)
		}
	}
	fun1()
	a4 *= 2
	fun1()
	fmt.Printf("--------------闭包的使用----------------------\n")

}

func swap(x, y string) (string, string) {
	return y, x
}
