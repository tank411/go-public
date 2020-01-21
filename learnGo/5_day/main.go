package main

//学习下多线程，创建，执行任务，销毁，
//网络io的了解
import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Println("第五天学习！")
	go say1()
	go say2()
	//var str string
	//fmt.Scan(&str)

}

func say1() {
	for index := 0; index < 10; index++ {
		fmt.Println("say1")
		log.Fatal("132")
		runtime.Gosched()
	}
}

func say2() {
	for index := 0; index < 10; index++ {
		fmt.Println("say2")
		runtime.Gosched()
	}
}
