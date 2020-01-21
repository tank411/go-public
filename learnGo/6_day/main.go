/*
 * @Descripttion: 模块信息介绍
 * @Company: HIK
 * @version: 1.0.0
 * @Author: TK
 * @Date: 2019-06-25 16:23:37
 * @LastEditors: TK
 * @LastEditTime: 2019-10-23 15:35:50
 */
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("json")

	st := &Student{
		"xiao ming",
		16,
	}

	b, err := json.Marshal(st)
	if err != nil {
		fmt.Println("encoding faild")
	} else {
		fmt.Println("encoded data : ")
		fmt.Println(b)
		fmt.Println(string(b))
	}

	ch := make(chan string, 1)
	go func(c chan string, str string) {
		c <- str
	}(ch, string(b))

	strData := <-ch
	stb := &Student{}

	err = json.Unmarshal([]byte(strData), &stb)

	if err != nil {
		fmt.Println("Unmarshal faild")
	} else {
		fmt.Println("Unmarshal success")
	}
	var stu []Student
	stu = append(stu, Student{"xiao ming", 16})
	stu = append(stu, Student{"x m", 15})

	js, _ := json.Marshal(stu)
	fmt.Println(js)

	var stu1 []Student
	json.Unmarshal(js, &stu1)
	fmt.Println(stu1)
}
