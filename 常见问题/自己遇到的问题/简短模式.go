package main

import "fmt"

var x = 100

func main() {
	fmt.Println(x, &x) //全局变量
	x := 1000
	fmt.Println(x, &x) //重新定义和初始化同名局部变量
	fmt.Println("---")
	selfTest01()
	fmt.Println("---")
	selfTest02()
	fmt.Println("---")
	selfTest03()
}

//退化赋值的前提条件是：至少有一个新变量被定义、且必须是同一作用域
func selfTest01() {
	x := 100
	fmt.Println(&x)
	x, y := 200, "abc"
	fmt.Println(&x, x)
	fmt.Println(y)
}

func selfTest02() {
	x := 100
	fmt.Println(x)
	//x := 200  no new variables on left side of :=
	//fmt.Println(x)
}

func selfTest03() {
	x := 100
	fmt.Println(x, &x)
	{
		x, y := 200, 300 //不同作用域、全部是新变量定义
		fmt.Println(&x, x, y)
	}
}
