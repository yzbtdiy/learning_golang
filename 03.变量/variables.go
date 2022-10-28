package main

import "fmt"

func main() {
	var a = "initial" //关键字var用于声明变量,声明同时可以赋值
	fmt.Println(a)

	var b, c int = 1, 2 //一次可以死声明多个变量
	fmt.Println(b, c)

	var d = true //go可以自动推断有初始值的变量类型
	fmt.Println(d)

	var e int      //变量声明后没有赋值,会初始化为零值
	fmt.Println(e) //int类型的零值为0

	f := "short" //:=是声明并初始化的简写,完整为var f string = "short"
	fmt.Println(f)
}

// ---result---
// initial
// 1 2
// true
// 0
// short
