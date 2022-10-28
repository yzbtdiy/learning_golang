package main //在go程序中,数组是一个具有编号且长度固定的元素序列,相较于数组,用的更多的是切片(slice)

import "fmt"

func main() {
	var a [5]int           //创建一个可以存放5个int元素的数组a,元素的类型和长度都是数组类型的一部分
	fmt.Println("emp:", a) //数组默认值为零值, int数组的元素零值是0

	a[4] = 100 //array[index]=value语法可以设置数组指定位置的值
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) //array[index]可以得到数组指定位置的值

	fmt.Println("len:", len(a)) //内置函数len可以返回数组的长度

	b := [5]int{1, 2, 3, 4, 5} //可以一行内声明并初始化数组
	fmt.Println("dcl:", b)

	var twoD [2][3]int //数组类型是一维的,但是可以组合构造多维数据结构
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)
}

// ---result---
// emp: [0 0 0 0 0]
// set: [0 0 0 0 100]
// get: 100
// len: 5
// dcl: [1 2 3 4 5]
// 2d: [[0 1 2] [1 2 3]]
