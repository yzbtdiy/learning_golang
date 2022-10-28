package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("write ", i, " as ")
	switch i { //switch是多分支情况时快捷的条件语句
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday: //同一个case中可以使用逗号来分隔多个表达式
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch { //不带表达式的switch是实现if else逻辑的另一种方式
	case t.Hour() < 12: //case表达式也可以不使用常量
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) { //类型开关(type switch)比较类型而非值,可以用来发现接口值的类型
		switch t := i.(type) { //变量t在每个分支中会有相应的类型
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}

// ---result---
// write 2 as two
// It's a weekday
// I'm a bool
// I'm an int
// Don't know type string
