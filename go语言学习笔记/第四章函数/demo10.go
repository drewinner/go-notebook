package 第四章函数

import "fmt"

func Test10() {
	x, y := 1, 2
	defer func(a int) {
		fmt.Println("defer x,y =", a, y)
	}(x)
	x += 100
	y += 200
	fmt.Println(x, y) //1,202
}

func Test1001() {
	x, y := 1, 2
	defer fmt.Println()
	x += 100
	y += 200
	fmt.Println(x, y) //1,202
}
