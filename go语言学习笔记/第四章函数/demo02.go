package 第四章函数

import "fmt"

func Test() *int {
	a := 1
	return &a
}

func Test02(x, y int, s string, _ bool) {
	fmt.Println(&x)
	x, z := 1, 2 //退化赋值
	fmt.Println(&x, &z)
}
