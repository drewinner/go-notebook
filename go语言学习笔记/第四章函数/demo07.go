package 第四章函数

import "fmt"

func TestStruct() {
	type calc struct {
		mul func(x, y int) int
	}
	x := calc{mul: func(x, y int) int {
		return x * y
	}}
	fmt.Println(x.mul(1, 2))
}

func TestChannel() {
	c := make(chan func(int, int) int, 2)
	c <- func(i int, i2 int) int {
		return i + i2
	}
	fmt.Println((<-c)(1, 2))
}
