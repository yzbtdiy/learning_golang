package main //声明文件所在的包,每个.go文件必须有归属的包

import "fmt" //为了使用包内函数,需要将其引入

func main() { //main主函数 程序入口
	fmt.Println("Hello,world") //控制台输出"Hello,world"
}

// ---result---
// Hello,world
