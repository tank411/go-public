package main

/*
文件的导入引用。
*/
import (
	"fmt"
	"learnGo/3_day/add"
	_ "unsafe"
)

/*
包含包中的私有变量或者私有函数的使用，必须使用
“//go:linkname num1 learnGo/3_day/add.num1 ”
将变量名或者函数名这样写出来，告诉编译器在链接时候找该位置的符号。
同事要导入 _ "unsafe"这个包
*/

//go:linkname num1 learnGo/3_day/add.num1
var num1 int

/*
补充说明：
Println :可以打印出字符串，和变量
Printf : 只可以打印出格式化的字符串,
可以输出字符串类型的变量，不可以输出整形变量和整形
*/
func main() {
	fmt.Printf("主函数开始！\n")
	result := add.Add(2, 3)
	fmt.Printf("相加结果为：%d\n", result)
	c := add.Num1
	fmt.Printf("num 的值为：%d\n", c)

	fmt.Printf("num 的值为：%d\n", num1)
	num1 = 20
	fmt.Printf("num 的值为：%d\n", num1)

	for index := 0; index < len(array1); index++ {
		fmt.Printf("array1 第%d的值为：%d\n", index, array1[index])
	}
	fmt.Printf("------------------------\n")
	//使用range关键字
	for i, num := range array1 {
		fmt.Printf("array1 第%d的值为：%d\n", i, num)
	}
	fmt.Printf("------------------------\n")
	//不需要使用该元素的序号，所以我们使用空白符"_"省略了
	for _, num := range array1 {
		fmt.Printf("值为：%d\n", num)
	}
	fmt.Printf("------------------------\n")
	for num := range array1 {
		fmt.Printf("值为：%d\n", num)
	}

	fmt.Printf("------------------------\n")
	sum := getArraySum(array1)
	fmt.Printf("sum 的值为：%d\n", sum)

	fmt.Printf("------------------------\n")
	fmt.Println(books{"标题", "名称"})
	fmt.Printf("------------------------\n")

	add.Call()

	add.Sub(2, 1)
}

//------------------------------------------------
//数组的认识
//格式 var variable_name[size] variable_type
//声明一个存放int16类型大小为10数组，
var array [10]int16

//初始化
var array1 = []int{1, 2, 3, 4, 5}

//说明：[...]int{1,2,3}代表存放了三个int类型的数组
//求数组长度: len(array1)
//数组作为参数

func getArraySum(arr []int) int {
	var sum int = 0
	for index := 0; index < len(arr); index++ {
		sum += arr[index]
	}
	return sum
}

//重点说明：未定义长度的数组做参数时候只能传递给未限制长度的函数，
//定义了长度的数组做参数，只能传递给限制长度的函数

//---------------------------------------------
var pArr *[]int = &array1 //c++写法 int *p[]；
//指针的认识
var ip *int //go的写法 c++的写法int *ip；

//--------------------------------------------
//结构体 type struct_variable_name struct{}
type books struct {
	title string
	name  string
}

//-----------------------------------------
//切片类似于动态数组，初始化方式很多。、
//s :=make([]int,len,cap) 通过make初始化切片s，类型为整形，len长度。=，cap容量。
//http://www.runoob.com/go/go-slice.html
