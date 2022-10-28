package main //range用于地带各种各样的数据结构
import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums { //使用range对slice中的元素求和,数组也可以通过这种方式初始化并赋初值
		sum += num //不需要索引可以使用空白标识符_将其忽略
	}
	fmt.Println("sum:", sum)

	for i, num := range nums { //range在数组和slice中提供对每项的索引和值的访问
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs { //range在map中迭代键值对
		fmt.Printf("%s -> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println("key:", k)
	}

	for i, c := range "go" { //range在字符串中迭代unicode码点(code point)
		fmt.Println(i, c) //第一个返回值是字符的起始字节位置,第二个是字符本身

	}
}

// ---result---
// sum: 9
// index: 1
// a -> apple
// b -> banana
// key: a
// key: b
// 0 103
// 1 111
