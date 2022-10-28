package main //slice是go中一个重要的数据类型,他提供了比数组更强大的交互方式

//https://go.dev/blog/slices-intro

import "fmt"

func main() { //与数组不同,slice的长度由它所包含的数组类型决定(与元素个数无关)
	s := make([]string, 3) //创建长度不为0的空slice需要使用内建函数make,此处创建了长度为3的string类型slice(初始值为零值)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s) //切片可以像数组一样设置和得到值
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s)) //len可以返回slice的长度

	s = append(s, "d")      //slice支持比数组更丰富的操作,比如
	s = append(s, "e", "f") //内建函数append会返回一个包含了一个或多个新值的slice,我们需要接受其返回值

	fmt.Println("apd:", s)

	c := make([]string, len(s)) //slice可以copy,需要创建一个空的和s有相同长度的c,然后将s复制给c
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]            //slice支持通过slice[low:high]语法进行"切片"操作(包含low但不包含high)
	fmt.Println("sl1:", l) //此处得到一个包含s[2],s[3]和s[4]的新slice

	l = s[:5]
	fmt.Println("sl2:", l) //新的slice包含从s[0]到s[5](不包括5)的元素

	l = s[2:]
	fmt.Println("sl3:", l) //新的slice包含从s[2](包含2)到结尾的元素

	t := []string{"g", "h", "i"} //可以在一行代码中声明并初始化slice
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3) //slice可以组成多维数组结构
	for i := 0; i < 3; i++ { //与多维数组不同的是内部slice长度可以不一样
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD) //slice和数组是不同的类型,但是通过fmt.Println输出结果是类似的
}

// ---result---
// emp: [  ]
// set: [a b c]
// get: c
// len: 3
// apd: [a b c d e f]
// cpy: [a b c d e f]
// sl1: [c d e]
// sl2: [a b c d e]
// sl3: [c d e f]
// dcl: [g h i]
// 2d:  [[0] [1 2] [2 3 4]]
