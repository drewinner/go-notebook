package 第四章函数

import "fmt"

func test(f func()) {
	f()
}
func test1() func(int, int) int {
	return func(i int, i2 int) int {
		return i + i2
	}
}

//赋值给变量
func Test06() {
	add := func(str string) {
		fmt.Println(str)
	}
	add("world")

	//作为参数
	test(func() {
		fmt.Println("as arguments ..")
	})
	//作为返回值
	add1 := test1()
	fmt.Println(add1(1, 2))

}
