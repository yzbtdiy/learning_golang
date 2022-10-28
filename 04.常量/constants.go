package main

import (
	"fmt"
	"math"
)

func main() {
	const s string = "constant" //const用于声明常量
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n //常数表达式可以执行任意精度的计算
	fmt.Println(d)

	fmt.Println(int64(d)) //数值型常量没有确定的类型,直到被给定类型,如显示类型转化

	fmt.Println(math.Sin(n)) //数字可以根据上下文需要(变量赋值,函数调用等)自动确定类型
}

// ---result---
// constant
// 6e+11
// 600000000000
// -0.28470407323754404
