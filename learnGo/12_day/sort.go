package main

import (
	"fmt"
	"sort"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type Persons []Person

func (p Persons) Len() int {
	return len(p)
}

func (p Persons) Less(i, j int) bool {
	return p[i].Age < p[j].Age
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	ps := Persons{}
	ps = append(ps, Person{
		"张三", 31,
	})
	ps = append(ps, Person{
		"李四", 23,
	})
	ps = append(ps, Person{
		"王五", 40,
	})

	sort.Sort(ps)
	fmt.Println(ps)

	arr := []int64{6, 1, 2, 5, 2, 9, 8}
	fmt.Println("arr", &arr)
	fmt.Println("arr[0]", &arr[0])
	fmt.Println("arr[2]", &arr[2])
	arr[0], arr[2] = arr[2], arr[0]
	fmt.Println("arr[0]", &arr[0])
	fmt.Println("arr[2]", &arr[2])
	fmt.Println(arr[0], arr[2])

	{
		var a = []int{1, 2, 3}
		fmt.Printf("%p\n", &a[0])
		b := a
		fmt.Printf("%p\n", &b[0]) // 问题1、为什么b的地址跟a不同，他们不是共享底层数组吗？
		b = append(b, 0)
		fmt.Printf("%p\n", &b[0]) // 问题2、为什么b的地址维持不变？append时b的长度不够应该重新分配了内存空间吧
		fmt.Println(a)
		fmt.Println(b)

	}

	{
		et := time.Now().Truncate(time.Hour * 24)
		tn := et.Unix() - int64(3600)
		bt := time.Unix(tn, 0)
		fmt.Println(bt)

	}
}
