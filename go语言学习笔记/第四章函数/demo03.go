package 第四章函数

import "fmt"

func Test03(x *int) {
	fmt.Printf("pointer: %p,target %v\n", &x, x) //输出形参x的地址
}
