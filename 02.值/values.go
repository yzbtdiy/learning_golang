package main

import "fmt"

func main() {
	fmt.Println("go" + "lang") //字符串可以通过+拼接

	fmt.Println("1+1 = ", 1+1)      //整数运算
	fmt.Println("7.0/3.0", 7.0/3.0) //浮点数运算

	fmt.Println(true && false) //逻辑与
	fmt.Println(true || false) //逻辑或
	fmt.Println(!true)         //逻辑非
}

// ---result---
// golang
// 1+1 =  2
// 7.0/3.0 2.3333333333333335
// false
// true
// false
