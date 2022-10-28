package main

import "fmt"

func main() { //go没有三目运算符,即使是基本的条件判断,依然需要使用完整的if语句
	if 7%2 == 0 { //基本if else用法,条件语句圆括号不是必需的,但花括号是必需的
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { //可以不用else只用if语句
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 { //条件语句前可以用声明语句,声明的变量可以在条件分支中使用
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
