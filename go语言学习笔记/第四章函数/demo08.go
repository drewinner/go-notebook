package 第四章函数

import "fmt"

//闭包
func Test08(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

func Test0801(x int) func() {
	fmt.Println(&x)
	return func() {
		fmt.Println(&x, x)
	}
}
