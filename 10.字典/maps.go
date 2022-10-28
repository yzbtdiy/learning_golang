package main //map是go内建的关联数据类型,在一些其他语言中也被称为哈希(hash)或者字典(dict)
import "fmt"

func main() {
	m := make(map[string]int) //要创建一个空map,需要使用内建函数make(map[key-type]val-type)

	m["k1"] = 7 //使用典型的name[key]=value来设置键值对
	m["k2"] = 13

	fmt.Println("map:", m) //使用fmt.Println打印一个map会输出所有的键值对

	v1 := m["k1"]
	fmt.Println("v1:", v1) //使用name[key]可以获取特定键的值

	fmt.Println("len:", len(m)) //内建函数len可以返回一个map的键值对数量

	delete(m, "k2") //内建函数delete可以从一个map中移除键值对
	fmt.Println("map:", m)

	_, prs := m["k2"] //从一个map中取值时,可以选择是否接收第二个返回值,该值表明map是否存在这个键
	//不需要值可以用空白标识符(blank_identifier)将其忽略
	fmt.Println("prs:", prs) //第二个返回值可以用来消除键不存在和键的值为零值产生的歧义,如0和""

	n := map[string]int{"foo": 1, "bar": 2} //一行代码声明并初始化一个新map
	fmt.Println("map:", n)                  //使用fmt.Println打印map时,是以map[k:v k:v]的格式输出的
}

// ---result---
// map: map[k1:7 k2:13]
// v1: 7
// len: 2
// map: map[k1:7]
// prs: false
// map: map[bar:2 foo:1]
