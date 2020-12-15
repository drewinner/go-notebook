package main

import "fmt"

//常量值必须是编译器可确定的字符、字符串、数字或者布尔值
func main() {
	const x = 1 //可以在函数代码块中定义常量、不曾使用的常量不会引发编译错误
	fmt.Println(x)

	const (
		a         = iota
		b float32 = iota
		c         = iota
	)
	fmt.Println(a, b, c)

}
