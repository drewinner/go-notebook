package main

import "fmt"

func main() {
	day6Test01()
}

//这道题考的是类型别名与类型定义的区别。
//
func day6Test01() {
	type MyInt1 int   //基于类型 int 创建了新类型 MyInt1
	type MyInt2 = int //创建了 int 的类型别名 MyInt2

	var i int = 0
	//var i1 MyInt1 = i //编译错误
	var i2 MyInt2 = i
	fmt.Println(i2)
	//fmt.Println(i1, i2)
}
