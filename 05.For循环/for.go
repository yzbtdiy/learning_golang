package main

import "fmt"

func main() { //for是go中唯一的循环结构
	i := 1
	for i <= 3 { //单个条件循环
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j < 9; j++ { //经典的初始/条件/后续for循环
		fmt.Println(j)
	}

	for { //不带条件的for循环将一直重复执行,直到在循环体内使用break或return跳出循环
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue //使用continue直接进入下一次循环
		}
		fmt.Println(n)
	}
}

// ---result---
// 1
// 2
// 3
// 7
// 8
// loop
// 1
// 3
// 5
